package Program

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func CreateDisk(disk InfoMKDisk) {
	fmt.Println("-------------------- Crear disco --------------------")
	if disk.Path == "" {
		fmt.Println("Error: Parámetro obligatorio Path vacío")
		return
	} else if disk.Size == 0 {
		fmt.Println("Error: Parámetro obligatorio Size vacío")
		return
	}
	path := disk.Path
	if _, err := os.Stat(path); err == nil {
		fmt.Println("El disco", path, "ya existe")
		return
	} else if err != nil {
		fmt.Println("Creando disco en:", path)
	}
	file, err := os.Create(path)
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
	var tamano int
	if strings.ToLower(disk.Unit) == "k" {
		tamano = disk.Size * 1024
	} else if strings.ToLower(disk.Unit) == "m" {
		tamano = disk.Size * 1024 * 1024
	} else if strings.ToLower(disk.Unit) == "" {
		tamano = disk.Size * 1024 * 1024
	} else {
		fmt.Println("Error: Parámetro Unit incorrecto")
		return
	}
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
	var fit string
	if strings.ToLower(disk.Fit) == "bf" {
		fit = "bf"
	} else if strings.ToLower(disk.Fit) == "wf" {
		fit = "wf"
	} else {
		fit = "ff"
	}

	var mbr MBR
	mbr.DskFit = []byte(fit)
	mbr.Tamano = []byte(strconv.Itoa(tamano))
	mbr.DskSignature = []byte(strconv.Itoa(rand.Intn(10000)))
	mbr.FechaCreacion = []byte(time.Now().Format(time.RFC850))
	for i := range mbr.Partition {
		mbr.Partition[i].inicializarParticiones(fit, tamano)
	}

	//Escritura del MBR

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

	fmt.Println("-------------------- Disco creado --------------------")
	fmt.Println("Fecha de creación:", string(mbr.FechaCreacion))
	fmt.Println("Signature:", string(mbr.DskSignature))
	fmt.Println("Size:", string(mbr.Tamano))
	fmt.Println("Fit:", string(mbr.DskFit))
}

func (p *Partition) inicializarParticiones(fit string, tam int) {
	p.Status = []byte("0")
	p.Size = []byte(strconv.Itoa(0))
	p.Fit = []byte(fit)
	p.Start = []byte(strconv.Itoa(tam))
	p.Name = []byte("")
}
