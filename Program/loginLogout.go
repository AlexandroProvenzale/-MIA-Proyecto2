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

var SesionActiva = new(Sesion)

func LoginS(logi InfoLogin) {
	if logi.User == "" {
		fmt.Println("Error: Parámetro obligatorio User vacío")
		return
	} else if logi.Pass == "" {
		fmt.Println("Error: Parámetro obligatorio Pass vacío")
		return
	} else if logi.Id == "" {
		fmt.Println("Error: Parámetro obligatorio Id vacío")
		return
	}
	if SesionActiva.Status {
		fmt.Println("Ya hay una sesión activa, haga un logout para iniciar sesión con otra cuenta")
		return
	}
	path, partName := obtenerParticion(logi.Id)
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
	fmt.Println("--------------------------LOGIN---------------------------")
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

	initInoArchivo := BytesToInt(sb.InodeStart) + BytesToInt(sb.InodeSize)
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
		initBloArchivo := BytesToInt(sb.BlockStart) + BytesToInt(sb.BlockSize)*block[i]
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
	for i := range registros {
		if registros[i] != "" {
			registroIndividual := strings.Split(registros[i], ",")
			if strings.ToUpper(registroIndividual[1]) == "U" {
				if registroIndividual[3] == logi.User && registroIndividual[4] == logi.Pass {
					if registroIndividual[0] == "0" {
						fmt.Println("El usuario no existe, fue eliminado con anterioridad.")
						return
					} else if logi.User == "root" {
						SesionActiva.Perm = 777
					} else {
						SesionActiva.Perm = 664
					}
					SesionActiva.User = logi.User
					SesionActiva.Id = strings.ToUpper(logi.Id)
					SesionActiva.Status = true
					SesionActiva.Group = registroIndividual[2]
					UID, _ := strconv.Atoi(registroIndividual[0])
					SesionActiva.UID = UID
					SesionActiva.Path = path
					SesionActiva.Name = partName
					break
				}
			}
		}
	}
	for i := range registros {
		if registros[i] != "" {
			registroIndividual := strings.Split(registros[i], ",")
			if strings.ToUpper(registroIndividual[1]) == "G" {
				if registroIndividual[2] == SesionActiva.Group {
					GID, _ := strconv.Atoi(registroIndividual[0])
					SesionActiva.GID = GID
					break
				}
			}
		}
	}
	if SesionActiva.Status {
		fmt.Println("El usuario", logi.User, "inició sesión correctamente!")
	} else {
		fmt.Println("ERROR: Usuario o contraseña incorrecto, compruebe el ID también.")
	}
}

func LogoutS() {
	if SesionActiva.Status {
		fmt.Println("Sesión del usuario ", SesionActiva.User, " cerrada exitosamente")
		SesionActiva.User = ""
		SesionActiva.Id = ""
		SesionActiva.Status = false
		SesionActiva.Group = ""
		SesionActiva.Path = ""
		SesionActiva.Name = ""
	} else {
		fmt.Println("ERROR: No hay ninguna sesión activa actualmente")
	}
}
