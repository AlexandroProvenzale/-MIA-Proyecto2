package Program

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

func Mkgroup(name string) {
	if !SesionActiva.Status {
		fmt.Println("ERROR: Necesita tener una sesión activa para utilizar el comando mkgroup.")
	} else if name == "" {
		fmt.Println("ERROR: Parámetro obligatorio Name vacío.")
		return
	} else if len(name) > 10 {
		fmt.Println("ERROR: Nombre del grupo no debe exceder los 10 caracteres.")
		return
	} else if SesionActiva.User != "root" {
		fmt.Println("ERROR: Usuario", SesionActiva.User, "sin persmisos para utilizar el comando mkgroup.")
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
	fmt.Println("-------------------------MKGROUP--------------------------")
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

	initInoArchivo := BytesToInt(sb.InodeStart) + BytesToInt(sb.InodeSize) + 100
	if _, err := file.Seek(int64(initInoArchivo), 0); err != nil { // Situamos el puntero en el inicio del inodo con archivo users
		log.Fatal(err)
		return
	}
	var inodoArchivo Inodo
	inoArchBytes := make([]byte, unsafe.Sizeof(inodoArchivo)+100)
	if _, err := file.Read(inoArchBytes); err != nil {
		log.Fatal(err)
		return
	}
	buff = bytes.NewBuffer(inoArchBytes)
	dec = gob.NewDecoder(buff)
	if err := dec.Decode(&inodoArchivo); err != nil { // Obtenemos la información del archivo users
		log.Fatal(err)
		return
	}

	// AQUÍ OBTENER EL ARREGLO BLOQUE DEL INODO Y RECORRERLO PARA VER CUÁNTOS APUNTADORES DIRECTOS USA EL ARCHIVO

	block := BytesToIntArray(inodoArchivo.Block)
	var archivoExtraido string
	for i := range block {
		if block[i] == -1 {
			break
		}
		initBloArchivo := BytesToInt(sb.BlockStart) + (BytesToInt(sb.BlockSize)+200)*(i+1)
		if _, err := file.Seek(int64(initBloArchivo), 0); err != nil {
			log.Fatal(err)
			return
		}
		var bloqueArchivo BloqueArchivo
		bloArchBytes := make([]byte, unsafe.Sizeof(bloqueArchivo)+200)
		if _, err := file.Read(bloArchBytes); err != nil {
			log.Fatal(err)
			return
		}
		buff = bytes.NewBuffer(bloArchBytes)
		dec = gob.NewDecoder(buff)
		if err := dec.Decode(&bloqueArchivo); err != nil { // Obtenemos la información del archivo users
			log.Fatal(err)
			return
		}
		archivoExtraido += string(bloqueArchivo.Content)
	}
	registros := strings.Split(archivoExtraido, "\n")
	GID := 1
	for i := range registros {
		if registros[i] != "" {
			registroIndividual := strings.Split(registros[i], ",")
			if strings.ToUpper(registroIndividual[1]) == "G" {
				GID += 1
				if registroIndividual[2] == name {
					fmt.Println("ERROR: El grupo", name, "ya existe.")
					return
				}
			}
		}
	}
	archivoExtraido += strconv.Itoa(GID) + ",G," + name + "\n"
	// VERIFICAR EL NUEVO TAMAÑO DEL ARCHIVO, SI OCUPA MÁS DE
	numBloquesArchivo := int(len(archivoExtraido)/64) + 1
	var posibleInt int

	for contador := 0; contador < numBloquesArchivo; contador++ {
		//VERIFICAR EN QUÉ APUNTADOR DIRECTO TOCA ESCRIBIR LA INFORMACIÓN DEL ARCHIVO
		block := BytesToIntArray(inodoArchivo.Block)
		block[contador] = BytesToInt(sb.FirstBlock)
		inodoArchivo.Block = IntArrayToBytes(block)
		bloqueArch := new(BloqueArchivo)
		var stringApuntador string
		for i := contador * 64; i < (contador+1)*64; i++ {
			if i < len(archivoExtraido) {
				if archivoExtraido != "" {
					stringApuntador += string(archivoExtraido[i])
				} else {
					if len(archivoExtraido) != 0 {
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
		initBloArchivo := BytesToInt(sb.BlockStart) + (BytesToInt(sb.BlockSize)+200)*(contador+1)
		EscribirBloqueArchivo(bloqueArch, initBloArchivo, file)
	}
	inodoArchivo.Size = IntToBytes(len(archivoExtraido))
	EscribirInodo(&inodoArchivo, initInoArchivo, file)
	fmt.Println("Grupo", name, "creado con éxito.")
}

func Mkuser(mk InfoMkuser) {
	if !SesionActiva.Status {
		fmt.Println("ERROR: Necesita tener una sesión activa para utilizar el comando mkuser.")
	} else if SesionActiva.User != "root" {
		fmt.Println("ERROR: Usuario", SesionActiva.User, "sin persmisos para utilizar el comando mkuser.")
		return
	} else if mk.User == "" {
		fmt.Println("Error: Parámetro obligatorio Usuario vacío.")
		return
	} else if mk.Pass == "" {
		fmt.Println("Error: Parámetro obligatorio Pwd vacío.")
		return
	} else if mk.Grp == "" {
		fmt.Println("Error: Parámetro obligatorio Grp vacío.")
		return
	} else if len(mk.User) > 10 {
		fmt.Println("ERROR: Nombre del usuario no debe exceder los 10 caracteres.")
		return
	} else if len(mk.Pass) > 10 {
		fmt.Println("ERROR: Contraseña del usuario no debe exceder los 10 caracteres.")
		return
	} else if len(mk.Grp) > 10 {
		fmt.Println("ERROR: Grupo del usuario no debe exceder los 10 caracteres.")
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
			fmt.Println("ERROR: Problema al cerrar el disco.")
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
	fmt.Println("--------------------------MKUSER--------------------------")
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

	initInoArchivo := BytesToInt(sb.InodeStart) + BytesToInt(sb.InodeSize) + 100
	if _, err := file.Seek(int64(initInoArchivo), 0); err != nil { // Situamos el puntero en el inicio del inodo con archivo users
		log.Fatal(err)
		return
	}
	var inodoArchivo Inodo
	inoArchBytes := make([]byte, unsafe.Sizeof(inodoArchivo)+100)
	if _, err := file.Read(inoArchBytes); err != nil {
		log.Fatal(err)
		return
	}
	buff = bytes.NewBuffer(inoArchBytes)
	dec = gob.NewDecoder(buff)
	if err := dec.Decode(&inodoArchivo); err != nil { // Obtenemos la información del archivo users
		log.Fatal(err)
		return
	}

	// AQUÍ OBTENER EL ARREGLO BLOQUE DEL INODO Y RECORRERLO PARA VER CUÁNTOS APUNTADORES DIRECTOS USA EL ARCHIVO

	block := BytesToIntArray(inodoArchivo.Block)
	var archivoExtraido string
	for i := range block {
		if block[i] == -1 {
			break
		}
		initBloArchivo := BytesToInt(sb.BlockStart) + (BytesToInt(sb.BlockSize)+200)*(i+1)
		if _, err := file.Seek(int64(initBloArchivo), 0); err != nil {
			log.Fatal(err)
			return
		}
		var bloqueArchivo BloqueArchivo
		bloArchBytes := make([]byte, unsafe.Sizeof(bloqueArchivo)+200)
		if _, err := file.Read(bloArchBytes); err != nil {
			log.Fatal(err)
			return
		}
		buff = bytes.NewBuffer(bloArchBytes)
		dec = gob.NewDecoder(buff)
		if err := dec.Decode(&bloqueArchivo); err != nil { // Obtenemos la información del archivo users
			log.Fatal(err)
			return
		}
		archivoExtraido += string(bloqueArchivo.Content)
	}
	registros := strings.Split(archivoExtraido, "\n")
	UID := 1
	existeGrupo := false
	for i := range registros {
		if registros[i] != "" {
			registroIndividual := strings.Split(registros[i], ",")
			if strings.ToUpper(registroIndividual[1]) == "U" {
				UID += 1
				if registroIndividual[3] == mk.User {
					fmt.Println("ERROR: El usuario", mk.User, "ya existe.")
					return
				}
			} else if strings.ToUpper(registroIndividual[1]) == "G" {
				if registroIndividual[2] == mk.Grp {
					existeGrupo = true
				}
			}
		}
	}
	if existeGrupo {
		archivoExtraido += strconv.Itoa(UID) + ",U," + mk.Grp + "," + mk.User + "," + mk.Pass + "\n"
	} else {
		fmt.Println("ERROR: El grupo", mk.Grp, "asignado al usuario", mk.User, "no existe.")
		return
	}
	// VERIFICAR EL NUEVO TAMAÑO DEL ARCHIVO, SI OCUPA MÁS DE
	numBloquesArchivo := int(len(archivoExtraido)/64) + 1
	var posibleInt int

	for contador := 0; contador < numBloquesArchivo; contador++ {
		//VERIFICAR EN QUÉ APUNTADOR DIRECTO TOCA ESCRIBIR LA INFORMACIÓN DEL ARCHIVO
		block := BytesToIntArray(inodoArchivo.Block)
		block[contador] = BytesToInt(sb.FirstBlock)
		inodoArchivo.Block = IntArrayToBytes(block)
		bloqueArch := new(BloqueArchivo)
		var stringApuntador string
		for i := contador * 64; i < (contador+1)*64; i++ {
			if i < len(archivoExtraido) {
				if archivoExtraido != "" {
					stringApuntador += string(archivoExtraido[i])
				} else {
					if len(archivoExtraido) != 0 {
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
		initBloArchivo := BytesToInt(sb.BlockStart) + (BytesToInt(sb.BlockSize)+200)*(contador+1)
		EscribirBloqueArchivo(bloqueArch, initBloArchivo, file)
	}
	inodoArchivo.Size = IntToBytes(len(archivoExtraido))
	EscribirInodo(&inodoArchivo, initInoArchivo, file)
	fmt.Println("Usuario", mk.User, "del grupo", mk.Grp, "creado con éxito")
}
