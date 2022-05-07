package Program

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"syscall"
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
	} else if strings.ToLower(rep.Name) == "file" {
		if rep.Ruta == "" {
			fmt.Println("Error: Parámetro obligatorio Ruta vacío, es necesario para este reporte")
			return
		}
		ReporteFile(path, partName, rep.Ruta)
	} else {
		ReportDisk(path, partName, rep.Path)
	}
}

func ReportDisk(path, partName, repo string) {
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

	var particion1 string
	var particion2 string
	var particion3 string
	var particion4 string
	var libre1 int
	var libre2 int
	var libre3 int
	var libre4 int
	var libre5 int
	var espacioLibre1 string
	var espacioLibre2 string
	var espacioLibre3 string
	var espacioLibre4 string
	var espacioLibre5 string

	var particionesOrden [4]Partition
	for i := 0; i < 4; i++ {
		particionesOrden[i] = mbr.Partition[i]
	}
	var auxiliar Partition
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if BytesToInt(particionesOrden[j].Start) > BytesToInt(particionesOrden[j+1].Start) {
				auxiliar = particionesOrden[j]
				particionesOrden[j] = particionesOrden[j+1]
				particionesOrden[j+1] = auxiliar
			}
		}
	}

	if BytesToInt(particionesOrden[0].Start) != BytesToInt(mbr.Tamano) {
		particion1 = "|" + string(particionesOrden[0].Name) + "\\n" + strconv.Itoa(BytesToInt(particionesOrden[0].Size)/BytesToInt(mbr.Tamano))
	}
	if BytesToInt(particionesOrden[1].Start) != BytesToInt(mbr.Tamano) {
		particion2 = "|" + string(particionesOrden[1].Name) + "\\n" + strconv.Itoa(BytesToInt(particionesOrden[1].Size)/BytesToInt(mbr.Tamano))
	}
	if BytesToInt(particionesOrden[2].Start) != BytesToInt(mbr.Tamano) {
		particion3 = "|" + string(particionesOrden[2].Name) + "\\n" + strconv.Itoa(BytesToInt(particionesOrden[2].Size)/BytesToInt(mbr.Tamano))
	}
	if BytesToInt(particionesOrden[3].Start) != BytesToInt(mbr.Tamano) {
		particion4 = "|" + string(particionesOrden[3].Name) + "\\n" + strconv.Itoa(BytesToInt(particionesOrden[3].Size)/BytesToInt(mbr.Tamano))
	}
	libre1 = BytesToInt(particionesOrden[0].Start) - int(unsafe.Sizeof(mbr))
	libre2 = BytesToInt(particionesOrden[1].Start) - (BytesToInt(particionesOrden[0].Start) + BytesToInt(particionesOrden[0].Start))
	libre3 = BytesToInt(particionesOrden[2].Start) - (BytesToInt(particionesOrden[1].Start) + BytesToInt(particionesOrden[1].Start))
	libre4 = BytesToInt(particionesOrden[3].Start) - (BytesToInt(particionesOrden[2].Start) + BytesToInt(particionesOrden[2].Start))
	libre5 = BytesToInt(mbr.Tamano) - (BytesToInt(particionesOrden[3].Start) + BytesToInt(particionesOrden[3].Size))

	if libre1 != 0 {
		espacioLibre1 = "|libre\\n" + strconv.Itoa(libre1/BytesToInt(mbr.Tamano))
	}
	if libre2 != 0 {
		espacioLibre1 = "|libre\\n" + strconv.Itoa(libre2/BytesToInt(mbr.Tamano))
	}
	if libre3 != 0 {
		espacioLibre1 = "|libre\\n" + strconv.Itoa(libre3/BytesToInt(mbr.Tamano))
	}
	if libre4 != 0 {
		espacioLibre1 = "|libre\\n" + strconv.Itoa(libre4/BytesToInt(mbr.Tamano))
	}
	if libre5 != 0 {
		espacioLibre1 = "|libre\\n" + strconv.Itoa(libre5/BytesToInt(mbr.Tamano))
	}

	inicioReporte := "digraph g {\n        rankdir = LR;\n    subgraph cluster0\n    {\n        "
	discoReporte := "Array [ shape = record, label = \"{"
	medioReporte := "MBR" + espacioLibre1 + particion1 + espacioLibre2 + particion2 + espacioLibre3 + particion3 +
		espacioLibre4 + particion4 + espacioLibre5
	finalReporte := "}\"] ;\n    }\n}"

	reporteCompleto := inicioReporte + discoReporte + medioReporte + finalReporte
	f, err := os.Create("reporteTree.dot")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	_, err2 := f.WriteString(reporteCompleto)
	if err2 != nil {
		log.Fatal(err2)
	}
	command := "dot -Tpng reporteDisk.dot -o \"" + repo + "\""
	System(command)
	fmt.Println("Reporte disk creado con éxito en", repo)
}

func ReporteFile(path, partName, ruta string) {
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
	fmt.Println("-----------------------REPORTE FILE-----------------------")
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
	rutas := strings.Split(ruta, "/")
	busquedaFile(&sb, rutas, initInode, &inodeActual, &blockActual, &pathCounter, BytesToInt(partition.Start), file, false, ruta)
}

func busquedaFile(sb *SuperBloque, rutas []string, initInode int, inodeActual, blockActual, pathCounter *int, sbStart int, file *os.File, existeRuta bool, fullruta string) {
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
							var contenidoArchivo string
							inodoArchivo := BytesToInt(sb.InodeStart) + BytesToInt(sb.InodeSize)*BytesToInt(bloqueCarpeta.BContent[j].Inodo)
							if _, err := file.Seek(int64(inodoArchivo), 0); err != nil {
								log.Fatal(err)
								return
							}
							var inArch Inodo
							inArchBytes := make([]byte, BytesToInt(sb.InodeSize))
							if _, err := file.Read(inArchBytes); err != nil {
								log.Fatal(err)
								return
							}
							buff = bytes.NewBuffer(inArchBytes)
							dec = gob.NewDecoder(buff)
							if err := dec.Decode(&inArch); err != nil {
								log.Fatal(err)
								return
							}
							bloqueArchivos := BytesToIntArray(inArch.Block)
							for x := range bloqueArchivos {
								if bloqueArchivos[x] == -1 {
									break
								}
								initBloArchivo := BytesToInt(sb.BlockStart) + BytesToInt(sb.BlockSize)*bloqueArchivos[x]
								if _, err := file.Seek(int64(initBloArchivo), 0); err != nil {
									log.Fatal(err)
									return
								}
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
								contenidoArchivo += string(bloqueArchivo.Content)
							}
							fmt.Println("Archivo:", rutas[*pathCounter])
							fmt.Println("Ruta completa:", fullruta)
							fmt.Println("Contiene lo siguiente:")
							fmt.Println("----------------------------------------------------------")
							fmt.Println(contenidoArchivo)
							fmt.Println("----------------------------------------------------------")
						}
						*pathCounter += 1
						// Hacer algo cuando encuentra la carpeta dentro de un bloque de carpetas
						*inodeActual = BytesToInt(bloqueCarpeta.BContent[j].Inodo)
						initIn := BytesToInt(sb.InodeStart) + BytesToInt(sb.InodeSize)**inodeActual
						busquedaFile(sb, rutas, initIn, inodeActual, blockActual, pathCounter, sbStart, file, existeRuta, fullruta)
						return
					}
				}
			}
		}
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
	f, err := os.Create("reporteTree.dot")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	_, err2 := f.WriteString(digraph)
	if err2 != nil {
		log.Fatal(err2)
	}
	command := "dot -Tpng reporteTree.dot -o \"" + repo + "\""
	System(command)
	fmt.Println("Reporte tree creado con éxito en", repo)
}

func System(cmd string) int {
	c := exec.Command("sh", "-c", cmd)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()

	if err == nil {
		return 0
	}

	// Figure out the exit code
	if ws, ok := c.ProcessState.Sys().(syscall.WaitStatus); ok {
		if ws.Exited() {
			return ws.ExitStatus()
		}

		if ws.Signaled() {
			return -int(ws.Signal())
		}
	}

	return -1
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
		"<TR>\n<TD>UID</TD>\n<TD>" + string(inodoActual.UID) + "</TD>\n</TR>\n" +
		"<TR>\n<TD>GID</TD>\n<TD>" + string(inodoActual.GID) + "</TD>\n</TR>\n" +
		"<TR>\n<TD>Type</TD>\n<TD>" + string(inodoActual.Type) + "</TD>\n</TR>\n" +
		"<TR>\n<TD>Size</TD>\n<TD>" + string(inodoActual.Size) + "</TD>\n</TR>\n" +
		"<TR>\n<TD>Perm</TD>\n<TD>" + string(inodoActual.Perm) + "</TD>\n</TR>\n" +
		"<TR>\n<TD>Mtime</TD>\n<TD>" + string(inodoActual.Mtime) + "</TD>\n</TR>\n" +
		"<TR>\n<TD>Ctime</TD>\n<TD>" + string(inodoActual.Ctime) + "</TD>\n</TR>\n"
	block := BytesToIntArray(inodoActual.Block)

	for i := range block {
		*digraph += "<TR>\n<TD PORT=\"ii" + strconv.Itoa(i) + "\">AD" + strconv.Itoa(i) + "</TD>\n<TD PORT=\"id" + strconv.Itoa(i) + "\">" + strconv.Itoa(block[i]) + "</TD>\n</TR>\n"
	}

	*digraph += "</TABLE>>];\n"

	for i := range block {
		if block[i] == -1 {
			break
		}

		*blockCounter = block[i]
		initBlo := BytesToInt(sb.BlockStart) + BytesToInt(sb.BlockSize)**blockCounter
		if _, err := file.Seek(int64(initBlo), 0); err != nil {
			log.Fatal(err)
			return
		}
		*digraph += "inode" + strconv.Itoa(*inodeCounter) + ":id" + strconv.Itoa(i) + " -> block" + strconv.Itoa(*blockCounter) + ";\n"
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
				"<TABLE BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"0\">\n" +
				"<TR>\n<TD COLSPAN=\"2\" BGCOLOR=\"#ff6363\">Bloque carpeta " + strconv.Itoa(*blockCounter) + "</TD>\n</TR>\n"

			for j := range bloqueCarpeta.BContent {
				if string(bloqueCarpeta.BContent[j].Name) != "" {
					*digraph += "<TR>\n<TD PORT=\"bi" + strconv.Itoa(j) + "\">" + string(bloqueCarpeta.BContent[j].Name) + "</TD>\n<TD PORT=\"bd" + strconv.Itoa(j) + "\">" + string(bloqueCarpeta.BContent[j].Inodo) + "</TD>\n</TR>\n"
				} else {
					*digraph += "<TR>\n<TD PORT=\"bi" + strconv.Itoa(j) + "\">-</TD>\n<TD PORT=\"bd" + strconv.Itoa(j) + "\">-1</TD>\n</TR>\n"
				}
			}

			*digraph += "</TABLE>>];\n"

			for j := range bloqueCarpeta.BContent {
				if string(bloqueCarpeta.BContent[j].Inodo) == "-1" {
					break
				} else if string(bloqueCarpeta.BContent[j].Name) != "." && string(bloqueCarpeta.BContent[j].Name) != ".." {
					//inodeCounter
					*inodeCounter = BytesToInt(bloqueCarpeta.BContent[j].Inodo)
					*blockCounter = block[i]
					*digraph += "block" + strconv.Itoa(*blockCounter) + ":bd" + strconv.Itoa(j) + " -> inode" + strconv.Itoa(*inodeCounter) + ";\n"
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
