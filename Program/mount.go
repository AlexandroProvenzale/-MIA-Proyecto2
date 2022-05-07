package Program

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"strconv"
	"unsafe"
)

var letras = [10]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
var DiscosM [5]Discos

func MontarDisco(disco InfoMount) {
	if disco.Path == "" {
		fmt.Println("Error: Parámetro obligatorio Path vacío")
		return
	} else if disco.Name == "" {
		fmt.Println("Error Parámetro obligatorio Name vacío")
		return
	}
	if _, err := os.Stat(disco.Path); err != nil {
		fmt.Println("Error: El disco en la ruta", disco.Path, "no existe")
		return
	}
	file, err := os.OpenFile(disco.Path, os.O_RDWR, 0777)
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
	var i int
	for i = 0; i < 5; i++ {
		if DiscosM[i].State == 0 || DiscosM[i].Path == disco.Path {
			break
		}
	}
	if i == 5 {
		fmt.Println("No hay espacio para montar el disco")
		return
	}
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
	partitionName := "noname"
	actualPartition := 5
	for j := 0; j < 4; j++ {
		partition := mbr.Partition[j]
		if string(partition.Status) == "1" && string(partition.Name) == disco.Name {
			fmt.Println("Partición:", j+1)
			fmt.Println("Estado:", string(partition.Status))
			fmt.Println("Tamaño:", string(partition.Size))
			fmt.Println("Fit:", string(partition.Fit))
			fmt.Println("Nombre:", string(partition.Name))
			fmt.Println("Start:", string(partition.Start))
			partitionName = string(partition.Name)
			actualPartition = j
			break
		}
	}
	if !ExistName(partitionName, mbr) {
		fmt.Println("No existe una partición con este nombre")
		return
	}
	DiscosM[i].State = 1
	DiscosM[i].Path = disco.Path
	DiscosM[i].Number = i + 1
	var contadorParticion int
	for x := 0; x < 10; x++ {
		mPartition := DiscosM[i].Partitions[x]
		if mPartition.State == 0 {
			DiscosM[i].Partitions[x].State = 1
			if actualPartition < 5 {
				DiscosM[i].Partitions[x].Name = string(mbr.Partition[actualPartition].Name)
			} else {
				fmt.Println("ERROR: Falló la ejecucuón del montaje")
			}
			DiscosM[i].Partitions[x].Letra = letras[x]
			DiscosM[i].Partitions[x].Identificador = "12" + strconv.Itoa(DiscosM[i].Number) + DiscosM[i].Partitions[x].Letra
			contadorParticion = x
			break
		} else {
			if mPartition.Name == string(mbr.Partition[actualPartition].Name) {
				fmt.Println("ERROR: La partición", string(mbr.Partition[actualPartition].Name), "ya fue montada")
				return
			}
		}
	}
	fmt.Println("El identificador de esta partición es:", DiscosM[i].Partitions[contadorParticion].Identificador)
	fmt.Println("Partición montada")
	fmt.Println("----------------------------------------------------------")
	fmt.Println("-Listado de particiones montadas-")
	for x := range DiscosM {
		if DiscosM[x].State == 1 {
			fmt.Println("Montadas en el disco", DiscosM[x].Path, ":")
			for j := range DiscosM[x].Partitions {
				if DiscosM[x].Partitions[j].State == 1 {
					fmt.Println("	*", DiscosM[x].Partitions[j].Name)
				}
			}
		}
	}
}
