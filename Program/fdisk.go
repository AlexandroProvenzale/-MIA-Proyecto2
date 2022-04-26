package Program

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

func CrearParticion(part InfoFDisk) {
	if part.Path == "" {
		fmt.Println("Error: Parámetro obligatorio Path vacío")
		return
	} else if part.Size <= 0 {
		fmt.Println("Error: Parámetro obligatorio Size vacío o menor a 0")
		return
	} else if part.Name == "" {
		fmt.Println("Error Parámetro obligatorio Name vacío")
		return
	}
	if _, err := os.Stat(part.Path); err != nil {
		fmt.Println("Error: El disco en la ruta", part.Path, "no existe")
		return
	}
	file, err := os.OpenFile(part.Path, os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error: Problema al cerrar el disco")
			return
		}
	}(file)
	if _, err := file.Seek(0, 0); err != nil {
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
	if err := dec.Decode(&mbr); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("-------------------- Disco encontrado --------------------")
	fmt.Println("Fecha de creación:", string(mbr.FechaCreacion))
	fmt.Println("Signature:", string(mbr.DskSignature))
	fmt.Println("Size:", string(mbr.Tamano))
	if string(mbr.DskFit) == "" {
		fmt.Println("Fit: ff")
	} else {
		fmt.Println("Fit:", string(mbr.DskFit))
	}
	fmt.Println("----------------------------------------------------------")

	var existent bool
	for i := range mbr.Partition {
		if string(mbr.Partition[i].Name) == part.Name {
			existent = true
		}
	}
	if existent {
		fmt.Println("Ya existe una particion con este nombre")
		return
	} else {
		var index int
		for i := range mbr.Partition {
			if string(mbr.Partition[i].Status) == "0" {
				index = i
				break
			}
		}
		if string(mbr.Partition[index].Status) != "0" {
			fmt.Println("----------------------------------------------------------")
			fmt.Println("ERROR: No hay espacio para otra partición, Ya ocupó 4")
			fmt.Println("----------------------------------------------------------")
		}

		var tamano int
		switch strings.ToLower(part.Unit) {
		case "b":
			tamano = part.Size
		case "m":
			tamano = part.Size * 1024 * 1024
		default:
			tamano = part.Size * 1024
		}

		switch strings.ToLower(part.Fit) {
		case "bf":
			mbr.Partition[index].Fit = []byte("bf")
		case "ff":
			mbr.Partition[index].Fit = []byte("ff")
		default:
			mbr.Partition[index].Fit = []byte("wf")
		}

		if strings.ToLower(part.Type) == "" || strings.ToLower(part.Type) == "p" || strings.ToLower(part.Type) == "e" {
			if strings.ToLower(part.Type) == "" || strings.ToLower(part.Type) == "p" {
				mbr.Partition[index].Size = []byte(strconv.Itoa(tamano))
				mbr.Partition[index].Status = []byte("1")
				mbr.Partition[index].Type = []byte("p")
			} else {
				if existsExtended, _ := existeExtendida(&mbr); existsExtended {
					fmt.Println("ERROR: Solo puede haber una partición extendida por disco")
					return
				}
				mbr.Partition[index].Size = []byte(strconv.Itoa(tamano))
				mbr.Partition[index].Status = []byte("1")
				mbr.Partition[index].Type = []byte("e")
			}

			if strings.ToLower(string(mbr.Partition[index].Fit)) == "ff" {
				fmt.Println("Creando partición con first fit")
				var trigger bool
				for i := 0; i < 4; i++ {
					for j := 0; j < 4; j++ {
						if !trigger {
							posibleInicial, _ := strconv.Atoi(string(mbr.Partition[getMinor(&mbr)].Start)) // Start de la primera partición del disco
							posibleInicial -= int(unsafe.Sizeof(mbr))                                      // Le resta el tamaño del mbr
							actualStart, _ := strconv.Atoi(string(mbr.Partition[i].Start))
							actualSize, _ := strconv.Atoi(string(mbr.Partition[i].Size))
							posible, _ := strconv.Atoi(string(mbr.Partition[getIndexFollow(&mbr, i)].Start)) // Comienzo de la partición siguiente
							posible = posible - (actualStart + actualSize)                                   // Le resta al comienzo de la partición siguiente, la posición final actual

							particionCreableSize, _ := strconv.Atoi(string(mbr.Partition[index].Size))
							if posibleInicial > particionCreableSize { // Significa que la partición creada es la primera en el disco, y que si cabe en el mismo
								mbr.Partition[index].Start = []byte(strconv.Itoa(int(unsafe.Sizeof(mbr)) + 1))
								trigger = true // Primera partición creada
								break          // Introducido en el first fit
							} else if posible > particionCreableSize { // Significa que la partición cabe en el disco
								mbr.Partition[index].Start = []byte(strconv.Itoa(actualStart + actualSize + 1))
								trigger = true // Partición creada
								break          // Introducido en el first fit
							}
						} else {
							break
						}
					}
				}
				if !trigger {
					fmt.Println("No hay espacio para crear la partición")
					mbr.Partition[index].Status = []byte("0")
					mbr.Partition[index].Size = []byte(strconv.Itoa(0))
					mbr.Partition[index].Fit = mbr.DskFit
					mbr.Partition[index].Start = mbr.Tamano
					mbr.Partition[index].Name = []byte("")
					return
				}
			} else if strings.ToLower(string(mbr.DskFit)) == "bf" {
				indexMejor := 12
				inicial, _ := strconv.Atoi(string(mbr.Partition[getMinor(&mbr)].Start))
				inicial -= int(unsafe.Sizeof(mbr))
				var definitivo int
				particionCreableSize, _ := strconv.Atoi(string(mbr.Partition[index].Size))
				if inicial > particionCreableSize {
					definitivo = inicial
				} else {
					definitivo = math.MaxInt32
				}
				for i := 0; i < 4; i++ {
					for j := 0; j < 4; j++ {
						StartJ, _ := strconv.Atoi(string(mbr.Partition[j].Start))
						StartI, _ := strconv.Atoi(string(mbr.Partition[i].Start))
						SizeI, _ := strconv.Atoi(string(mbr.Partition[i].Size))
						posible := StartJ - (StartI + SizeI)
						if posible < definitivo && posible > particionCreableSize {
							indexMejor = i
							definitivo = posible
						}
					}
				}
				if definitivo != math.MaxInt32 {
					if indexMejor == 12 {
						mbr.Partition[index].Start = []byte(strconv.Itoa(int(unsafe.Sizeof(mbr)) + 1))
					} else {
						Startim, _ := strconv.Atoi(string(mbr.Partition[indexMejor].Start))
						Sizeim, _ := strconv.Atoi(string(mbr.Partition[indexMejor].Size))
						mbr.Partition[index].Start = []byte(strconv.Itoa(Startim + Sizeim + 1))
					}
				} else {
					fmt.Println("No hay espacio para crear la partición")
					mbr.Partition[index].Status = []byte("0")
					mbr.Partition[index].Size = []byte(strconv.Itoa(0))
					mbr.Partition[index].Fit = mbr.DskFit
					mbr.Partition[index].Start = mbr.Tamano
					mbr.Partition[index].Name = []byte("")
					return
				}
			} else {
				indexMejor := 12
				inicial, _ := strconv.Atoi(string(mbr.Partition[getMinor(&mbr)].Start))
				inicial -= int(unsafe.Sizeof(mbr))
				var definitivo int
				particionCreableSize, _ := strconv.Atoi(string(mbr.Partition[index].Size))
				if inicial > particionCreableSize {
					definitivo = inicial
				} else {
					definitivo = math.MinInt32
				}
				for i := 0; i < 4; i++ {
					for j := 0; j < 4; j++ {
						StartS, _ := strconv.Atoi(string(mbr.Partition[getIndexFollow(&mbr, i)].Start))
						StartI, _ := strconv.Atoi(string(mbr.Partition[i].Start))
						SizeI, _ := strconv.Atoi(string(mbr.Partition[i].Size))
						posible := StartS - (StartI + SizeI)
						if posible > definitivo && posible > particionCreableSize {
							indexMejor = i
							definitivo = posible
						}
					}
				}
				if definitivo != math.MinInt32 {
					if indexMejor == 12 {
						mbr.Partition[index].Start = []byte(strconv.Itoa(int(unsafe.Sizeof(mbr)) + 1))
					} else {
						Startim, _ := strconv.Atoi(string(mbr.Partition[indexMejor].Start))
						Sizeim, _ := strconv.Atoi(string(mbr.Partition[indexMejor].Size))
						mbr.Partition[index].Start = []byte(strconv.Itoa(Startim + Sizeim + 1))
					}
				} else {
					fmt.Println("No hay espacio para crear la partición")
					mbr.Partition[index].Status = []byte("0")
					mbr.Partition[index].Size = []byte(strconv.Itoa(0))
					mbr.Partition[index].Fit = mbr.DskFit
					mbr.Partition[index].Start = mbr.Tamano
					mbr.Partition[index].Name = []byte("")
					return
				}
			}
		} else {
			existsExtended, indexExtended := existeExtendida(&mbr)
			if !existsExtended {
				fmt.Println("ERROR: Debe haber una partición extendida en la que crear una partición lógica")
				return
			}
			var logicPartition EBR
			logicPartition.Size = []byte(strconv.Itoa(part.Size))
			logicPartition.Name = []byte(part.Name)
			posicionExtendida, _ := strconv.Atoi(string(mbr.Partition[indexExtended].Start))
			posicionExtendida += 1
			if _, err := file.Seek(int64(posicionExtendida), 0); err != nil {
				log.Fatal(err)
				return
			}
			var auxiliar EBR
			auxiliarbytes := make([]byte, unsafe.Sizeof(auxiliar))
			if _, err := file.Read(auxiliarbytes); err != nil {
				log.Fatal(err)
				return
			}
			buff = bytes.NewBuffer(auxiliarbytes)
			dec = gob.NewDecoder(buff)
			if err := dec.Decode(&auxiliar); err != nil {
				if err.Error() != "EOF" {
					log.Fatal(err)
					return
				}
			}
			var startExistente int
			auxNext, _ := strconv.Atoi(string(auxiliar.Next))
			auxStart, _ := strconv.Atoi(string(auxiliar.Start))
			auxSize, _ := strconv.Atoi(string(auxiliar.Size))
			for auxNext == (auxStart + auxSize + 1) {
				startExistente = auxStart
				if _, err := file.Seek(int64(auxNext), 0); err != nil {
					log.Fatal(err)
					return
				}
				auxiliarbytes = make([]byte, unsafe.Sizeof(auxiliar))
				if _, err := file.Read(auxiliarbytes); err != nil {
					log.Fatal(err)
					return
				}
				buff = bytes.NewBuffer(auxiliarbytes)
				dec = gob.NewDecoder(buff)
				if err := dec.Decode(&auxiliar); err != nil {
					if err.Error() != "EOF" {
						log.Fatal(err)
						return
					} else {
						break
					}
				}
				auxNext1, _ := strconv.Atoi(string(auxiliar.Next))
				auxNext = auxNext1
				auxStart1, _ := strconv.Atoi(string(auxiliar.Start))
				auxStart = auxStart1
				auxSize1, _ := strconv.Atoi(string(auxiliar.Size))
				auxSize = auxSize1
			}
			if startExistente != 0 {
				if _, err := file.Seek(int64(startExistente), 0); err != nil {
					log.Fatal(err)
					return
				}
				auxiliarbytes = make([]byte, unsafe.Sizeof(auxiliar))
				if _, err := file.Read(auxiliarbytes); err != nil {
					log.Fatal(err)
					return
				}
				buff = bytes.NewBuffer(auxiliarbytes)
				dec = gob.NewDecoder(buff)
				if err := dec.Decode(&auxiliar); err != nil {
					if err.Error() != "EOF" {
						log.Fatal(err)
						return
					}
				}
				logicPartition.Start = auxiliar.Next
				logicPartitionStart, _ := strconv.Atoi(string(logicPartition.Start))
				logicPartitionSize, _ := strconv.Atoi(string(logicPartition.Size))
				logicPartition.Next = []byte(strconv.Itoa(logicPartitionStart + logicPartitionSize + 1))
			} else {
				extendedPartitionStart, _ := strconv.Atoi(string(mbr.Partition[indexExtended].Start))
				extendedPartitionStart += 1
				logicPartition.Start = []byte(strconv.Itoa(extendedPartitionStart))
				logicPartitionStart, _ := strconv.Atoi(string(logicPartition.Start))
				logicPartitionSize, _ := strconv.Atoi(string(logicPartition.Size))
				logicPartition.Next = []byte(strconv.Itoa(logicPartitionStart + logicPartitionSize + 1))
			}
			logicPartitionStart, _ := strconv.Atoi(string(logicPartition.Start))
			if _, err := file.Seek(int64(logicPartitionStart), 0); err != nil {
				log.Fatal(err)
				return
			}
			buff = new(bytes.Buffer)
			enc := gob.NewEncoder(buff)
			if err := enc.Encode(logicPartition); err != nil {
				log.Fatal(err)
				return
			}
			if _, err := file.Write(buff.Bytes()); err != nil {
				log.Fatal(err)
				return
			}
			fmt.Println("La partición lógica", part.Name, "comienza en", string(logicPartition.Start))
			fmt.Println("Partición creada con éxito")
		}
		if strings.ToLower(part.Type) == "" || strings.ToLower(part.Type) == "e" || strings.ToLower(part.Type) == "p" {
			mbr.Partition[index].Name = []byte(part.Name)
			fmt.Println("La partición", part.Name, "comienza en", string(mbr.Partition[index].Start))
			fmt.Println("Partición creada con éxito")
		}
	}
	if _, err := file.Seek(0, 0); err != nil {
		log.Fatal(err)
		return
	}
	buff = new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	if err := enc.Encode(mbr); err != nil {
		log.Fatal(err)
		return
	}
	if _, err := file.Write(buff.Bytes()); err != nil {
		log.Fatal(err)
		return
	}
}

func getMinor(mbr *MBR) int {
	indexMinor := 0
	for i := 0; i < 4; i++ {
		Start1, _ := strconv.Atoi(string(mbr.Partition[i].Start))
		Start2, _ := strconv.Atoi(string(mbr.Partition[indexMinor].Start))
		if Start1 < Start2 {
			indexMinor = i
		}
	}
	return indexMinor // DEVUELVE LA PARTICIÓN QUE TIENE EL MENOR START
}

func getIndexFollow(mbr *MBR, index int) int { // Index es la partición actual en el ciclo for i
	indexReturn := index
	sizeBetween := math.MaxInt32
	Start, _ := strconv.Atoi(string(mbr.Partition[index].Start)) // Start de la partición actual en el ciclo for i
	Size, _ := strconv.Atoi(string(mbr.Partition[index].Size))   // Size de la partición actual en el ciclo for i
	finalIndex := Start + Size                                   // Posición final de la partición actual en el ciclo for i
	for i := 0; i < 4; i++ {
		Start1, _ := strconv.Atoi(string(mbr.Partition[i].Start))     // Vuelve a verificar el Start de cada partición
		if finalIndex < Start1 && (Start1-finalIndex) < sizeBetween { // Si la posición final de la partición actual en el for i es menor al comienzo de la partición en este ciclo temporal
			sizeBetween = Start1 - finalIndex // Y la posición start del actual ciclo temporal menos el final de la partición actual for i es menor a sizeBetween
			indexReturn = i                   // Actualiza el sizeBetween y el index que va a retornar
		}
	}
	return indexReturn // retorna el index de la partición que le sigue a la partición actual del for i
}

func existeExtendida(mbr *MBR) (bool, int) {
	for i := range mbr.Partition {
		if strings.ToLower(string(mbr.Partition[i].Type)) == "e" {
			return true, i
		}
	}
	return false, -1
}
