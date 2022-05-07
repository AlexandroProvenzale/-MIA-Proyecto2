package Program

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"unsafe"
)

func EscribirDirectorio(dir InfoMkdir) {
	if !SesionActiva.Status {
		fmt.Println("ERROR: Necesita tener una sesión activa para utilizar el comando mkgroup.")
	} else if dir.Path == "" {
		fmt.Println("ERROR: Parámetro obligatorio Path vacío.")
		return
	}
	path := SesionActiva.Path
	partName := SesionActiva.Name
	if _, err := os.Stat(path); err != nil { // Verificamos que la ruta esté correctamente almacenada
		fmt.Println("ERROR: El disco en la ruta", path, "no existe.")
		return
	}
	file, err := os.OpenFile(path, os.O_RDWR, 0777) // Abrimos el archivo
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(file *os.File) { // Posponemos su cierre hasta que el método termine su ejecución
		err := file.Close()
		if err != nil {
			fmt.Println("Error: Problema al cerrar el disco.")
			return
		}
	}(file)

	if _, err := file.Seek(0, 0); err != nil { // Situamos el puntero en el principio
		log.Fatal(err)
		return
	}
	var mbr MBR
	mbrbytes := make([]byte, unsafe.Sizeof(mbr))
	if _, err := file.Read(mbrbytes); err != nil {
		log.Fatal(err)
		return
	}
	buff := bytes.NewBuffer(mbrbytes)
	dec := gob.NewDecoder(buff)
	if err := dec.Decode(&mbr); err != nil { // Obtenemos la información del mbr para buscar la partición
		log.Fatal(err)
		return
	}
	fmt.Println("---------------------------MKDIR--------------------------")
	var partition Partition
	for j := 0; j < 4; j++ { // Buscamos la partición con el nombre almacenado que respecta a la del ID ingresado
		partition = mbr.Partition[j]
		if string(partition.Status) == "1" && string(partition.Name) == partName {
			break
		}
	}
	if reflect.DeepEqual(partition, Partition{}) {
		fmt.Println("No se encontró la partición", partName, "en el disco", path)
		return
	}

	if _, err := file.Seek(int64(BytesToInt(partition.Start)), 0); err != nil { // Situamos el puntero en el inicio de la particion
		log.Fatal(err)
		return
	}
	var sb SuperBloque
	sbBytes := make([]byte, unsafe.Sizeof(sb)+50)
	if _, err := file.Read(sbBytes); err != nil {
		log.Fatal(err)
		return
	}
	buff = bytes.NewBuffer(sbBytes)
	dec = gob.NewDecoder(buff)
	if err := dec.Decode(&sb); err != nil { // Obtenemos la información del superbloque
		log.Fatal(err)
		return
	}
	inodeActual := 0
	blockActual := 0
	pathCounter := 0
	initInode := BytesToInt(sb.InodeStart)
	rutas := strings.Split(dir.Path, "/")
	buscandoRecursivamente(&sb, rutas, initInode, &inodeActual, &blockActual, &pathCounter, dir.P, BytesToInt(partition.Start), file, false, dir.Path)
}

func buscandoRecursivamente(sb *SuperBloque, rutas []string, initInode int, inodeActual, blockActual, pathCounter *int, permiso bool, sbStart int, file *os.File, existeRuta bool, nombreRuta string) {
	if _, err := file.Seek(int64(initInode), 0); err != nil {
		log.Fatal(err)
		return
	}
	var inodoActual Inodo
	inodeBytes := make([]byte, BytesToInt(sb.InodeSize))
	if _, err := file.Read(inodeBytes); err != nil {
		log.Fatal(err)
		return
	}
	buff := bytes.NewBuffer(inodeBytes)
	dec := gob.NewDecoder(buff)
	if err := dec.Decode(&inodoActual); err != nil {
		log.Fatal(err)
		return
	}
	block := BytesToIntArray(inodoActual.Block)
	for i := range block {
		if block[i] == -1 {
			// something
			break
		}
		*blockActual = block[i]
		if *blockActual == 0 && rutas[*pathCounter] == "" {
			*pathCounter += 1
		}
		initBlo := BytesToInt(sb.BlockStart) + BytesToInt(sb.BlockSize)**blockActual
		if _, err := file.Seek(int64(initBlo), 0); err != nil {
			log.Fatal(err)
			return
		}
		if BytesToInt(inodoActual.Type) == 0 {
			var bloqueCarpeta BloqueCarpeta
			bloCarpBytes := make([]byte, BytesToInt(sb.BlockSize))
			if _, err := file.Read(bloCarpBytes); err != nil {
				log.Fatal(err)
				return
			}
			buff = bytes.NewBuffer(bloCarpBytes)
			dec = gob.NewDecoder(buff)
			if err := dec.Decode(&bloqueCarpeta); err != nil {
				log.Fatal(err)
				return
			}

			for j := range bloqueCarpeta.BContent {
				if string(bloqueCarpeta.BContent[j].Inodo) == "-1" {
					// something
					break
				} else if string(bloqueCarpeta.BContent[j].Name) != "." && string(bloqueCarpeta.BContent[j].Name) != ".." {
					//inodeActual
					if string(bloqueCarpeta.BContent[j].Name) == rutas[*pathCounter] {
						if rutas[*pathCounter] == rutas[len(rutas)-1] {
							fmt.Println("Error", rutas[*pathCounter], "ya existe en esta carpeta")
							return
						}
						*pathCounter += 1
						// Hacer algo cuando encuentra la carpeta dentro de un bloque de carpetas
						*inodeActual = BytesToInt(bloqueCarpeta.BContent[j].Inodo)
						initIn := BytesToInt(sb.InodeStart) + BytesToInt(sb.InodeSize)**inodeActual
						buscandoRecursivamente(sb, rutas, initIn, inodeActual, blockActual, pathCounter, permiso, sbStart, file, existeRuta, nombreRuta)
						return
					}
				}
			}
			if !existeRuta {
				if rutas[*pathCounter] == rutas[len(rutas)-1] {
					for j := range bloqueCarpeta.BContent {
						if string(bloqueCarpeta.BContent[j].Name) == rutas[*pathCounter] {
							fmt.Println("ERROR", rutas[*pathCounter], "ya existe")
							return
						}
						if string(bloqueCarpeta.BContent[j].Inodo) == "-1" {
							bloqueCarpeta.BContent[j].Name = []byte(rutas[*pathCounter])
							bloqueCarpeta.BContent[j].Inodo = sb.FirstInode
							EscribirBloqueCarpeta(&bloqueCarpeta, initBlo, file)
							initIno := BytesToInt(sb.InodeStart) + BytesToInt(sb.InodeSize)**inodeActual
							EscribirInodo(&inodoActual, initIno, file)
							CrearCarpeta(sb, inodeActual, file)
							EscribirSuperBloque(sb, sbStart, file)
							existeRuta = true
							break
						}
					}
				} else {
					if permiso {
						for j := range bloqueCarpeta.BContent {
							if string(bloqueCarpeta.BContent[j].Inodo) == "-1" {
								bloqueCarpeta.BContent[j].Name = []byte(rutas[*pathCounter])
								bloqueCarpeta.BContent[j].Inodo = sb.FirstInode
								EscribirBloqueCarpeta(&bloqueCarpeta, initBlo, file)
								initIno := BytesToInt(sb.InodeStart) + BytesToInt(sb.InodeSize)**inodeActual
								EscribirInodo(&inodoActual, initIno, file)
								CrearCarpeta(sb, inodeActual, file)
								EscribirSuperBloque(sb, sbStart, file)
								*inodeActual = BytesToInt(bloqueCarpeta.BContent[j].Inodo)
								*pathCounter += 1
								initIn := BytesToInt(sb.InodeStart) + BytesToInt(sb.InodeSize)**inodeActual
								buscandoRecursivamente(sb, rutas, initIn, inodeActual, blockActual, pathCounter, permiso, sbStart, file, existeRuta, nombreRuta)
								return
							}
						}
					} else {
						fmt.Println("ERROR: La ruta ingresada no existe, hace falta parámetro -P para crearla")
						return
					}
				}
			}
		}
	}
	if !existeRuta {
		if rutas[*pathCounter] == rutas[len(rutas)-1] {
			// Si la ruta es la última, osea la que se está creando
			for i := range block {

				if block[i] == -1 {
					block[i] = BytesToInt(sb.FirstBlock)
					inodoActual.Block = IntArrayToBytes(block)
					initIno := BytesToInt(sb.InodeStart) + BytesToInt(sb.InodeSize)**inodeActual
					EscribirInodo(&inodoActual, initIno, file)
					bloqueApuntador := new(BloqueCarpeta)
					bloqueApuntador.BContent[0].Name = []byte(rutas[*pathCounter])
					bloqueApuntador.BContent[1].Name = []byte("")
					bloqueApuntador.BContent[2].Name = []byte("")
					bloqueApuntador.BContent[3].Name = []byte("")
					bloqueApuntador.BContent[0].Inodo = sb.FirstInode
					bloqueApuntador.BContent[1].Inodo = IntToBytes(-1)
					bloqueApuntador.BContent[2].Inodo = IntToBytes(-1)
					bloqueApuntador.BContent[3].Inodo = IntToBytes(-1)
					EscribirBloqueCarpeta(bloqueApuntador, GetInitBlock(sb), file)
					ActualizarBitmap(sb.BmBlockStart, &sb.FirstBlock, &sb.FreeBlocksCount, BytesToInt(sb.BlocksCount), file)
					CrearCarpeta(sb, inodeActual, file)
					EscribirSuperBloque(sb, sbStart, file)
					existeRuta = true
					break
				}
			}
		} else {
			if permiso {
				for i := range block {
					if block[i] == -1 {
						block[i] = BytesToInt(sb.FirstBlock)
						inodoActual.Block = IntArrayToBytes(block)
						initIno := BytesToInt(sb.InodeStart) + BytesToInt(sb.InodeSize)**inodeActual
						EscribirInodo(&inodoActual, initIno, file)
						bloqueApuntador := new(BloqueCarpeta)
						bloqueApuntador.BContent[0].Name = []byte(rutas[*pathCounter])
						bloqueApuntador.BContent[1].Name = []byte("")
						bloqueApuntador.BContent[2].Name = []byte("")
						bloqueApuntador.BContent[3].Name = []byte("")
						bloqueApuntador.BContent[0].Inodo = sb.FirstInode
						bloqueApuntador.BContent[1].Inodo = IntToBytes(-1)
						bloqueApuntador.BContent[2].Inodo = IntToBytes(-1)
						bloqueApuntador.BContent[3].Inodo = IntToBytes(-1)
						EscribirBloqueCarpeta(bloqueApuntador, GetInitBlock(sb), file)
						ActualizarBitmap(sb.BmBlockStart, &sb.FirstBlock, &sb.FreeBlocksCount, BytesToInt(sb.BlocksCount), file)
						CrearCarpeta(sb, inodeActual, file)
						EscribirSuperBloque(sb, sbStart, file)
						*inodeActual = BytesToInt(bloqueApuntador.BContent[0].Inodo)
						*pathCounter += 1
						initIn := BytesToInt(sb.InodeStart) + BytesToInt(sb.InodeSize)**inodeActual
						buscandoRecursivamente(sb, rutas, initIn, inodeActual, blockActual, pathCounter, permiso, sbStart, file, existeRuta, nombreRuta)
						return
					}
				}

			} else {
				fmt.Println("ERROR: La ruta ingresada no existe, hace falta parámetro -P para crearla")
				return
			}
		}
	}
	fmt.Println("Carpeta", nombreRuta, "creada exitosamente")
}

func CrearCarpeta(sb *SuperBloque, inodeActual *int, file *os.File) {
	newInodeFile := NewInodo(SesionActiva.UID, SesionActiva.GID, 0, 0, SesionActiva.Perm)
	blockIno := BytesToIntArray(newInodeFile.Block)
	blockIno[0] = BytesToInt(sb.FirstBlock)
	newInodeFile.Block = IntArrayToBytes(blockIno)
	newBlockFile := new(BloqueCarpeta)
	newBlockFile.BContent[0].Name = []byte("..")
	newBlockFile.BContent[1].Name = []byte(".")
	newBlockFile.BContent[2].Name = []byte("")
	newBlockFile.BContent[3].Name = []byte("")
	newBlockFile.BContent[0].Inodo = IntToBytes(*inodeActual)
	newBlockFile.BContent[1].Inodo = sb.FirstInode
	newBlockFile.BContent[2].Inodo = IntToBytes(-1)
	newBlockFile.BContent[3].Inodo = IntToBytes(-1)

	EscribirInodo(newInodeFile, GetInitInode(sb), file)
	ActualizarBitmap(sb.BmInodeStart, &sb.FirstInode, &sb.FreeInodesCount, BytesToInt(sb.InodesCount), file)
	EscribirBloqueCarpeta(newBlockFile, GetInitBlock(sb), file)
	ActualizarBitmap(sb.BmBlockStart, &sb.FirstBlock, &sb.FreeBlocksCount, BytesToInt(sb.BlocksCount), file)
}
