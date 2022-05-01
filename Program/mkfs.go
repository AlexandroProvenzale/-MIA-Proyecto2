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

func Formatear(form InfoMkfs) {
	if form.Id == "" {
		fmt.Println("Error: Parámetro obligatorio Id vacío")
		return
	}
	var path, partName string
	for esp := range DiscosM { // Verificar cada disco
		for part := range DiscosM[esp].Partitions { // Verificar cada particion montada por disco en busca del Id
			if DiscosM[esp].Partitions[part].Identificador == strings.ToUpper(form.Id) { // Si encuentra la partición montada
				path = DiscosM[esp].Path                      // Guarda el path
				partName = DiscosM[esp].Partitions[part].Name // Y guarda el nombre de la partición
				break
			}
		}
		if path != "" && partName != "" {
			break
		}
	}
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
			break
		}
	}
	if reflect.DeepEqual(partition, Partition{}) {
		fmt.Println("No se encontró la partición", partName, "en el disco", path)
		return
	}
	n := int(numeroEstructuras(partition))
	inicioParticion, _ := strconv.Atoi(string(partition.Start))
	var SBnuevo SuperBloque
	SBnuevo.FilesystemType = []byte("2")
	SBnuevo.InodesCount = []byte(strconv.Itoa(n))
	SBnuevo.BlocksCount = []byte(strconv.Itoa(n * 3))
	SBnuevo.FreeInodesCount = []byte(strconv.Itoa(n))
	SBnuevo.FreeBlocksCount = []byte(strconv.Itoa(n * 3))
	SBnuevo.Mtime = []byte(time.Now().Format(time.RFC850))
	SBnuevo.MntCount = []byte(strconv.Itoa(1))
	SBnuevo.Magic = []byte("0xEF53")
	SBnuevo.InodeSize = []byte(strconv.Itoa(int(unsafe.Sizeof(Inodo{}))))
	SBnuevo.BlockSize = []byte(strconv.Itoa(int(unsafe.Sizeof(BloqueArchivo{}))))
	SBnuevo.FirstInode = []byte(strconv.Itoa(0))
	SBnuevo.FirstBlock = []byte(strconv.Itoa(0))
	bmInodeStart := inicioParticion + int(unsafe.Sizeof(SuperBloque{})) + 1
	SBnuevo.BmInodeStart = []byte(strconv.Itoa(bmInodeStart))
	bmBlockStart := bmInodeStart + n + 1
	SBnuevo.BmBlockStart = []byte(strconv.Itoa(bmBlockStart))
	inodeStart := bmBlockStart + n*3 + 1
	SBnuevo.InodeStart = []byte(strconv.Itoa(inodeStart))
	blockStart := inodeStart + n*int(unsafe.Sizeof(Inodo{})) + 1
	SBnuevo.BlockStart = []byte(strconv.Itoa(blockStart))

	if form.Type == "full" || form.Type == "" {
		if _, err := file.Seek(int64(inicioParticion), 0); err != nil { // Situamos el puntero en el inicio de la partición
			log.Fatal(err)
			return
		}
		tamano, _ := strconv.Atoi(string(partition.Size))
		var b = make([]byte, tamano)
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
	EscribirBitmaps(bmInodeStart, n, file)
	EscribirBitmaps(bmBlockStart, n*3, file)
	HacerRoot(&SBnuevo, file)
	if _, err := file.Seek(int64(inicioParticion), 0); err != nil { // Situamos el puntero en el inicio de la partición
		log.Fatal(err)
		return
	}
	buff = new(bytes.Buffer)

	enc := gob.NewEncoder(buff)
	if err := enc.Encode(SBnuevo); err != nil {
		log.Fatal(err)
		return
	}
	if _, err := file.Write(buff.Bytes()); err != nil {
		log.Fatal(err)
		return
	}
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
	firstBloqueCarpeta.BContent[0].Name = []byte("..")
	firstBloqueCarpeta.BContent[1].Name = []byte(".")
	EscribirInodo(firstInode, GetInitInode(sb), file)
	ActualizarBitmap(sb.BmInodeStart, &sb.FirstInode, &sb.FreeInodesCount, BytesToInt(sb.InodesCount), file)
	EscribirBloqueCarpeta(firstBloqueCarpeta, GetInitBlock(sb), file)
	ActualizarBitmap(sb.BmBlockStart, &sb.FirstBlock, &sb.FreeBlocksCount, BytesToInt(sb.BlocksCount), file)
	// Crear archivo de usuarios
	CrearUsuarios(sb, file)
}

func CrearUsuarios(sb *SuperBloque, file *os.File) {
	users := "1,G,root\n1,U,root,root,123,\n"
	init := BytesToInt(sb.BlockStart)
	if _, err := file.Seek(int64(init), 0); err != nil { // Situamos el puntero en el final de la última estructura inodo
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
	if err := dec.Decode(&bloqueArchivosActual); err != nil { // Obtenemos la información del mbr para buscar la partición
		log.Fatal(err)
		return
	}
	bloqueArchivosActual.BContent[2].Inodo = sb.FirstBlock
	bloqueArchivosActual.BContent[2].Name = []byte("users.txt")
	EscribirBloqueCarpeta(&bloqueArchivosActual, init, file)
	// ¿Puedo escribir el archivo en este bloque y en este espacio?
	// Si se comprueba que se puede se procede con el método que identifica cuántos bloques de archivos deben de crearse (Cada uno de 64 bytes) y también lo crea de una vez
	CrearArchivo(sb, users, 1, 1, 777)
}

func CrearArchivo(sb *SuperBloque, archivo string, UID, GID, perm int) {
	tamArchivo := unsafe.Sizeof(archivo)
	numBloquesArchivo := int(unsafe.Sizeof(tamArchivo)/64) + 1
	inodoArchivo := NewInodo(UID, GID, int(unsafe.Sizeof(tamArchivo)), 1, perm)
	block := BytesToIntArray(inodoArchivo.Block)
	block[0] = BytesToInt(sb.FirstBlock)
	inodoArchivo.Block = IntArrayToBytes(block)
	for contador := 0; contador < numBloquesArchivo; contador++ {
		//VERIFICAR EN QUÉ APUNTADOR DIRECTO TOCA ESCRIBIR LA INFORMACIÓN DEL ARCHIVO
		
	}
}

func GetInitInode(sb *SuperBloque) int {
	return BytesToInt(sb.InodeStart) + BytesToInt(sb.InodeSize)*BytesToInt(sb.FirstInode)
}

func GetInitBlock(sb *SuperBloque) int {
	return BytesToInt(sb.BlockStart) + BytesToInt(sb.BlockSize)*BytesToInt(sb.FirstBlock)
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
		if bmBytes[i] == 0 {
			bmBytes[i] = 1
			break
		}
		if i+1 == size {
			fmt.Println("BITMAP LLENO")
			return
		}
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
