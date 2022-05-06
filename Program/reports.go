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

func Reportar(rep InfoRep) {
	if rep.Id == "" {
		fmt.Println("Error: Parámetro obligatorio Id vacío")
		return
	} else if rep.Name == "" {
		fmt.Println("Error: Parámetro obligatorio Name vacío")
		return
	} else if rep.Path == "" {
		fmt.Println("Error: Parámetro obligatorio Path vacío")
		return
	}
	path, partName := obtenerParticion(rep.Id)
	if _, err := os.Stat(path); err != nil { // Verificamos que la ruta esté correctamente almacenada
		fmt.Println("ERROR: El disco en la ruta", path, "no existe")
		return
	}

	if strings.ToLower(rep.Name) == "tree" {
		ReporteTree(path, partName, rep.Path)
	}
}

func ReporteTree(path, partName, repo string) {
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
	fmt.Println("-----------------------REPORTE TREE-----------------------")
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

	digraph := "digraph g {\nfontname=\"Helvetica,Arial,sans-serif\"\n" +
		"node [fontname=\"Helvetica,Arial,sans-serif\"]\n" +
		"edge [fontname=\"Helvetica,Arial,sans-serif\"]\n" +
		"graph [fontsize=30 labelloc=\"t\" label=\"\" splines=true overlap=false rankdir = \"LR\"];\n" +
		"node [\n" +
		"shape = \"plaintext\"\n" +
		"fontsize = \"18\"\n" +
		"style=\"filled\"\n" +
		"];\n"

	inodeCounter := 0
	blockCounter := -1
	initInode := BytesToInt(sb.InodeStart)
	recursivamente(&sb, &digraph, initInode, &inodeCounter, &blockCounter, file)
	digraph += "}"
	fmt.Println(digraph)
}

func recursivamente(sb *SuperBloque, digraph *string, initInode int, inodeCounter, blockCounter *int, file *os.File) {
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
	if err := dec.Decode(&inodoActual); err != nil { // Obtenemos la información del inodo
		log.Fatal(err)
		return
	}

	*digraph += "inode" + strconv.Itoa(*inodeCounter) + "[label=<\n" +
		"<TABLE BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"0\" CELLPADDING=\"4\">\n" +
		"<TR>\n<TD COLSPAN=\"2\" BGCOLOR=\"cornflowerblue\">Inodo " + strconv.Itoa(*inodeCounter) + "</TD>\n</TR>\n" +
		"<TR>\n<TD>Type</TD>\n<TD>" + string(inodoActual.Type) + "</TD>\n</TR>\n" +
		"<TR>\n<TD>Size</TD>\n<TD>" + string(inodoActual.Size) + "</TD>\n</TR>\n" +
		"<TR>\n<TD>Perm</TD>\n<TD>" + string(inodoActual.Perm) + "</TD>\n</TR>\n"
	block := BytesToIntArray(inodoActual.Block)

	for i := range block {
		*digraph += "<TR>\n<TD>AD" + strconv.Itoa(i) + "</TD>\n<TD PORT=\"i" + strconv.Itoa(i) + "\">" + strconv.Itoa(block[i]) + "</TD>\n</TR>\n"
	}

	*digraph += "</TABLE>>];\n"

	for i := range block {
		if block[i] == -1 {
			break
		}

		*blockCounter += 1
		initBlo := BytesToInt(sb.BlockStart) + BytesToInt(sb.BlockSize)**blockCounter
		if _, err := file.Seek(int64(initBlo), 0); err != nil {
			log.Fatal(err)
			return
		}
		*digraph += "inode" + strconv.Itoa(*inodeCounter) + ":i" + strconv.Itoa(i) + " -> block" + strconv.Itoa(*blockCounter) + ":b" + strconv.Itoa(block[i]) + "\n"
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

			*digraph += "block" + strconv.Itoa(*blockCounter) + "[label=<\n" +
				"<TABLE BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"0\" CELLPADDING=\"4\">\n" +
				"<TR>\n<TD COLSPAN=\"2\" BGCOLOR=\"#ff6363\">Bloque carpeta " + strconv.Itoa(*blockCounter) + "</TD>\n</TR>\n"

			for i := range bloqueCarpeta.BContent {
				if string(bloqueCarpeta.BContent[i].Name) != "" {
					*digraph += "<TR>\n<TD>" + string(bloqueCarpeta.BContent[i].Name) + "</TD>\n<TD PORT=\"b" + strconv.Itoa(i) + "\">" + string(bloqueCarpeta.BContent[i].Inodo) + "</TD>\n</TR>\n"
				} else {
					*digraph += "<TR>\n<TD>-</TD>\n<TD PORT=\"b" + strconv.Itoa(i) + "\">-1</TD>\n</TR>\n"
				}
			}

			*digraph += "</TABLE>>];\n"

			for i := range bloqueCarpeta.BContent {
				if string(bloqueCarpeta.BContent[i].Inodo) == "-1" {
					break
				} else if string(bloqueCarpeta.BContent[i].Name) != "." && string(bloqueCarpeta.BContent[i].Name) != ".." {
					//inodeCounter
					*inodeCounter += 1
					*digraph += "block" + strconv.Itoa(*blockCounter) + ":b" + strconv.Itoa(i) + " -> inode" + strconv.Itoa(*inodeCounter) + ":i" + string(bloqueCarpeta.BContent[i].Inodo) + "\n"
					initIno := BytesToInt(sb.InodeStart) + BytesToInt(sb.InodeSize)**inodeCounter
					recursivamente(sb, digraph, initIno, inodeCounter, blockCounter, file)
				}
			}
		} else if BytesToInt(inodoActual.Type) == 1 {
			var bloqueArchivo BloqueArchivo
			bloArchBytes := make([]byte, BytesToInt(sb.BlockSize))
			if _, err := file.Read(bloArchBytes); err != nil {
				log.Fatal(err)
				return
			}
			buff = bytes.NewBuffer(bloArchBytes)
			dec = gob.NewDecoder(buff)
			if err := dec.Decode(&bloqueArchivo); err != nil {
				log.Fatal(err)
				return
			}

			*digraph += "block" + strconv.Itoa(*blockCounter) + "[label=<\n" +
				"<TABLE BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"0\" CELLPADDING=\"4\">\n" +
				"<TR>\n<TD COLSPAN=\"2\" BGCOLOR=\"darkolivegreen2\">Bloque archivo " + strconv.Itoa(*blockCounter) + "</TD>\n</TR>\n"

			*digraph += "<TR>\n<TD COLSPAN=\"2\">" + string(bloqueArchivo.Content) + "</TD>\n</TR>\n"

			*digraph += "</TABLE>>];\n"
		}
	}
}
