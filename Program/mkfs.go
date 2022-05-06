package Program

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func obtenerParticion(Id string) (string, string) {
	var path, partName string
	for esp := range DiscosM { // Verificar cada disco
		for part := range DiscosM[esp].Partitions { // Verificar cada particion montada por disco en busca del Id
			if DiscosM[esp].Partitions[part].Identificador == strings.ToUpper(Id) { // Si encuentra la partición montada
				path = DiscosM[esp].Path                      // Guarda el path
				partName = DiscosM[esp].Partitions[part].Name // Y guarda el nombre de la partición
				break
			}
		}
		if path != "" && partName != "" {
			break
		}
	}
	return path, partName
}

func Formatear(form InfoMkfs) {
	if form.Id == "" {
		fmt.Println("Error: Parámetro obligatorio Id vacío")
		return
	}
	path, partName := obtenerParticion(form.Id)
	if _, err := os.Stat(path); err != nil { // Verificamos que la ruta esté correctamente almacenada
		fmt.Println("ERROR: El disco en la ruta", path, "no existe")
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
			fmt.Println("Error: Problema al cerrar el disco")
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
	fmt.Println("---------------------------MKFS---------------------------")
	fmt.Println("-------------------- Disco encontrado --------------------")
	fmt.Println("Fecha de creación:", string(mbr.FechaCreacion))
	fmt.Println("Signature:", string(mbr.DskSignature))
	fmt.Println("Size:", string(mbr.Tamano))
	fmt.Println("----------------------------------------------------------")
	var partition Partition
	for j := 0; j < 4; j++ { // Buscamos la partición con el nombre almacenado que respecta a la del ID ingresado
		partition = mbr.Partition[j]
		if string(partition.Status) == "1" && string(partition.Name) == partName {
			fmt.Println("Partición:", j+1, "del disco, encontrada")
			fmt.Println("Estado:", string(partition.Status))
			fmt.Println("Tamaño:", string(partition.Size))
			fmt.Println("Fit:", string(partition.Fit))
			fmt.Println("Nombre:", string(partition.Name))
			fmt.Println("Start:", string(partition.Start))
			fmt.Println("----------------------------------------------------------")
			break
		}
	}
	if reflect.DeepEqual(partition, Partition{}) {
		fmt.Println("No se encontró la partición", partName, "en el disco", path)
		return
	}
	n := int(numeroEstructuras(partition))
	inicioParticion := BytesToInt(partition.Start)

	var SBnuevo SuperBloque
	SBnuevo.FilesystemType = []byte("2")
	SBnuevo.InodesCount = IntToBytes(n)
	SBnuevo.BlocksCount = IntToBytes(n * 3)
	SBnuevo.FreeInodesCount = IntToBytes(n)
	SBnuevo.FreeBlocksCount = IntToBytes(n * 3)
	SBnuevo.Mtime = []byte(time.Now().Format(time.RFC850))
	SBnuevo.MntCount = IntToBytes(1)
	SBnuevo.Magic = []byte("0xEF53")
	SBnuevo.InodeSize = IntToBytes(int(unsafe.Sizeof(Inodo{})) + 100)
	SBnuevo.BlockSize = IntToBytes(int(unsafe.Sizeof(BloqueArchivo{})) + 200)
	SBnuevo.FirstInode = IntToBytes(0)
	SBnuevo.FirstBlock = IntToBytes(0)
	bmInodeStart := inicioParticion + int(unsafe.Sizeof(SuperBloque{})) + 50
	SBnuevo.BmInodeStart = IntToBytes(bmInodeStart)
	bmBlockStart := bmInodeStart + n + 1
	SBnuevo.BmBlockStart = IntToBytes(bmBlockStart)
	inodeStart := bmBlockStart + n*3 + 1
	SBnuevo.InodeStart = IntToBytes(inodeStart)
	blockStart := inodeStart + n*int(unsafe.Sizeof(Inodo{})) + 1
	SBnuevo.BlockStart = IntToBytes(blockStart)

	if form.Type == "full" || form.Type == "" {
		if _, err := file.Seek(int64(inicioParticion), 0); err != nil { // Situamos el puntero en el inicio de la partición
			log.Fatal(err)
			return
		}
		var b = make([]byte, BytesToInt(partition.Size))
		for i := range b {
			b[i] = 0
		}
		buff := new(bytes.Buffer)
		if err := binary.Write(buff, binary.LittleEndian, b); err != nil { // Convirtiendo arreglo de 0s en binario dentro del Buffer
			log.Fatal(err)
			return
		}
		if _, err := file.Write(buff.Bytes()); err != nil { // Escribiendo el Buffer dentro del archivo
			log.Fatal(err)
			return
		}
	}

	EscribirBitmaps(bmInodeStart, n, file)   // +
	EscribirBitmaps(bmBlockStart, n*3, file) // +
	HacerRoot(&SBnuevo, file)                // +
	EscribirSuperBloque(&SBnuevo, inicioParticion, file)
	fmt.Println("FORMATEO COMPLETO")
	fmt.Println("Filesystem type:", string(SBnuevo.FilesystemType))
	fmt.Println("Inodes count:", string(SBnuevo.InodesCount))
	fmt.Println("Blocks count:", string(SBnuevo.BlocksCount))
	fmt.Println("Free inodes count:", string(SBnuevo.FreeInodesCount))
	fmt.Println("Free blocks count:", string(SBnuevo.FreeBlocksCount))
	fmt.Println("mtime:", string(SBnuevo.Mtime))
	fmt.Println("mnt count:", string(SBnuevo.MntCount))
	fmt.Println("magic:", string(SBnuevo.Magic))
	fmt.Println("Inode size:", string(SBnuevo.InodeSize))
	fmt.Println("Block size:", string(SBnuevo.BlockSize))
	fmt.Println("First inode:", string(SBnuevo.FirstInode))
	fmt.Println("First block:", string(SBnuevo.FirstBlock))
	fmt.Println("BM inode start:", string(SBnuevo.BmInodeStart))
	fmt.Println("BM block start:", string(SBnuevo.BmBlockStart))
	fmt.Println("Inodes start:", string(SBnuevo.InodeStart))
	fmt.Println("Block start:", string(SBnuevo.BlockStart))
}

func numeroEstructuras(partition Partition) float64 {
	tamParticion, _ := strconv.Atoi(string(partition.Size))
	soSB := int(unsafe.Sizeof(SuperBloque{}))
	soInodo := int(unsafe.Sizeof(Inodo{}))
	soBloque := int(unsafe.Sizeof(BloqueArchivo{}))
	resultado := (tamParticion - soSB) / (4 + soInodo + 3*soBloque)
	return math.Floor(float64(resultado))
}

func HacerRoot(sb *SuperBloque, file *os.File) {
	firstInode := NewInodo(1, 1, 0, 0, 777)
	block := BytesToIntArray(firstInode.Block)
	block[0] = BytesToInt(sb.FirstBlock)
	firstInode.Block = IntArrayToBytes(block)
	firstBloqueCarpeta := new(BloqueCarpeta)
	firstBloqueCarpeta.BContent[0].Inodo = IntToBytes(0)
	firstBloqueCarpeta.BContent[1].Inodo = IntToBytes(0)
	firstBloqueCarpeta.BContent[2].Inodo = IntToBytes(-1)
	firstBloqueCarpeta.BContent[3].Inodo = IntToBytes(-1)
	firstBloqueCarpeta.BContent[0].Name = []byte("..")
	firstBloqueCarpeta.BContent[1].Name = []byte(".")
	firstBloqueCarpeta.BContent[2].Name = []byte("")
	firstBloqueCarpeta.BContent[3].Name = []byte("")
	EscribirInodo(firstInode, GetInitInode(sb), file)
	ActualizarBitmap(sb.BmInodeStart, &sb.FirstInode, &sb.FreeInodesCount, BytesToInt(sb.InodesCount), file)
	EscribirBloqueCarpeta(firstBloqueCarpeta, GetInitBlock(sb), file)
	ActualizarBitmap(sb.BmBlockStart, &sb.FirstBlock, &sb.FreeBlocksCount, BytesToInt(sb.BlocksCount), file)
	// Crear archivo de usuarios
	CrearUsuarios(sb, file)
}

func CrearUsuarios(sb *SuperBloque, file *os.File) {
	users := "1,G,root\n1,U,root,root,123\n"
	init := BytesToInt(sb.BlockStart)
	if _, err := file.Seek(int64(init), 0); err != nil { // Situamos el puntero en el inicio de las estructuras bloque
		log.Fatal(err)
		return
	}
	var bloqueArchivosActual BloqueCarpeta
	bdirBytes := make([]byte, unsafe.Sizeof(bloqueArchivosActual))
	if _, err := file.Read(bdirBytes); err != nil {
		log.Fatal(err)
		return
	}
	buff := bytes.NewBuffer(bdirBytes)
	dec := gob.NewDecoder(buff)
	if err := dec.Decode(&bloqueArchivosActual); err != nil {
		log.Fatal(err)
		return
	}
	bloqueArchivosActual.BContent[2].Inodo = sb.FirstInode
	bloqueArchivosActual.BContent[2].Name = []byte("users.txt")
	EscribirBloqueCarpeta(&bloqueArchivosActual, init, file)
	// ¿Puedo escribir el archivo en este bloque y en este espacio?
	// Si se comprueba que se puede se procede con el método que identifica cuántos bloques de archivos deben de crearse (Cada uno de 64 bytes) y también lo crea de una vez
	CrearArchivo(sb, users, len(users), 1, 1, 777, file)
	fmt.Println("Archivo users.txt creado en carpeta root ")
}

func CrearArchivo(sb *SuperBloque, archivo string, tamArchivo, UID, GID, perm int, file *os.File) {
	numBloquesArchivo := int(tamArchivo/64) + 1
	inodoArchivo := NewInodo(UID, GID, tamArchivo, 1, perm)
	var posibleInt int

	for contador := 0; contador < numBloquesArchivo; contador++ {
		//VERIFICAR EN QUÉ APUNTADOR DIRECTO TOCA ESCRIBIR LA INFORMACIÓN DEL ARCHIVO
		block := BytesToIntArray(inodoArchivo.Block)
		block[contador] = BytesToInt(sb.FirstBlock)
		inodoArchivo.Block = IntArrayToBytes(block)
		bloqueArch := new(BloqueArchivo)
		var stringApuntador string
		for i := contador * 64; i < (contador+1)*64; i++ {
			if i < tamArchivo {
				if archivo != "" {
					stringApuntador += string(archivo[i])
				} else {
					if tamArchivo != 0 {
						stringApuntador += strconv.Itoa(posibleInt)
						posibleInt += 1
						if posibleInt == 10 {
							posibleInt = 0
						}
					}
				}
			} else {
				break
			}
		}
		bloqueArch.Content = []byte(stringApuntador)
		EscribirBloqueArchivo(bloqueArch, GetInitBlock(sb), file)
		ActualizarBitmap(sb.BmBlockStart, &sb.FirstBlock, &sb.FreeBlocksCount, BytesToInt(sb.BlocksCount), file)
	}
	EscribirInodo(inodoArchivo, GetInitInode(sb), file)
	ActualizarBitmap(sb.BmInodeStart, &sb.FirstInode, &sb.FreeInodesCount, BytesToInt(sb.InodesCount), file)
}

func GetInitInode(sb *SuperBloque) int {
	return BytesToInt(sb.InodeStart) + BytesToInt(sb.InodeSize)*BytesToInt(sb.FirstInode)
}

func GetInitBlock(sb *SuperBloque) int {
	return BytesToInt(sb.BlockStart) + BytesToInt(sb.BlockSize)*BytesToInt(sb.FirstBlock)
}

func EscribirSuperBloque(super *SuperBloque, init int, file *os.File) {
	if _, err := file.Seek(int64(init), 0); err != nil { // Situamos el puntero en el inicio de la partición
		log.Fatal(err)
		return
	}
	buff := new(bytes.Buffer)

	enc := gob.NewEncoder(buff)
	if err := enc.Encode(super); err != nil {
		log.Fatal(err)
		return
	}
	if _, err := file.Write(buff.Bytes()); err != nil {
		log.Fatal(err)
		return
	}
}

func EscribirInodo(inodo *Inodo, init int, file *os.File) {
	if _, err := file.Seek(int64(init), 0); err != nil { // Situamos el puntero en el final de la última estructura inodo
		log.Fatal(err)
		return
	}
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	if err := enc.Encode(inodo); err != nil {
		log.Fatal(err)
		return
	}
	if _, err := file.Write(buff.Bytes()); err != nil {
		log.Fatal(err)
		return
	}
}

func EscribirBloqueCarpeta(bloque *BloqueCarpeta, init int, file *os.File) {
	if _, err := file.Seek(int64(init), 0); err != nil { // Situamos el puntero en el final de la última estructura inodo
		log.Fatal(err)
		return
	}
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	if err := enc.Encode(bloque); err != nil {
		log.Fatal(err)
		return
	}
	if _, err := file.Write(buff.Bytes()); err != nil {
		log.Fatal(err)
		return
	}
}

func EscribirBloqueArchivo(bloque *BloqueArchivo, init int, file *os.File) {
	if _, err := file.Seek(int64(init), 0); err != nil { // Situamos el puntero en el final de la última estructura inodo
		log.Fatal(err)
		return
	}
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	if err := enc.Encode(bloque); err != nil {
		log.Fatal(err)
		return
	}
	if _, err := file.Write(buff.Bytes()); err != nil {
		log.Fatal(err)
		return
	}
}

func ActualizarBitmap(start []byte, first, free *[]byte, size int, file *os.File) {
	if _, err := file.Seek(int64(BytesToInt(start)), 0); err != nil { // Situamos el puntero en al final del super bloque
		log.Fatal(err)
		return
	}
	bmBytes := make([]byte, size)
	if _, err := file.Read(bmBytes); err != nil {
		log.Fatal(err)
		return
	}
	for i := range bmBytes {
		if bmBytes[i] == 48 {
			bmBytes[i] = 49
			break
		}
		if i+1 == size {
			fmt.Println("BITMAP LLENO")
			return
		}
	}
	if _, err := file.Seek(int64(BytesToInt(start)), 0); err != nil { // Situamos el puntero en al final del super bloque
		log.Fatal(err)
		return
	}
	buff := new(bytes.Buffer)
	if err := binary.Write(buff, binary.LittleEndian, bmBytes); err != nil { // Convirtiendo arreglo de 0s en binario dentro del Buffer
		log.Fatal(err)
		return
	}
	if _, err := file.Write(buff.Bytes()); err != nil { // Escribiendo el Buffer dentro del archivo
		log.Fatal(err)
		return
	}
	*free = IntToBytes(BytesToInt(*free) - 1)
	*first = IntToBytes(BytesToInt(*first) + 1)
}

func EscribirBitmaps(start, sizeBM int, file *os.File) {
	if _, err := file.Seek(int64(start), 0); err != nil { // Situamos el puntero en al final del super bloque
		log.Fatal(err)
		return
	}
	var b = make([]byte, sizeBM)
	for i := range b {
		b[i] = 48
	}
	buff := new(bytes.Buffer)
	if err := binary.Write(buff, binary.LittleEndian, b); err != nil { // Convirtiendo arreglo de 0s en binario dentro del Buffer
		log.Fatal(err)
		return
	}
	if _, err := file.Write(buff.Bytes()); err != nil { // Escribiendo el Buffer dentro del archivo
		log.Fatal(err)
		return
	}
}
