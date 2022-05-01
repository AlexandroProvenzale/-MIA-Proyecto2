package Program

import (
	"bytes"
	"encoding/gob"
	"log"
	"strconv"
	"time"
)

type InfoMKDisk struct {
	Path string
	Fit  string
	Size int
	Unit string
}

type InfoFDisk struct {
	Path string
	Fit  string
	Size int
	Unit string
	Type string
	Name string
}

type InfoMount struct {
	Path string
	Name string
}

type InfoMkfs struct {
	Id   string
	Type string
}

type MBR struct {
	Tamano        []byte
	FechaCreacion []byte
	DskSignature  []byte
	DskFit        []byte
	Partition     [4]Partition
}

type EBR struct {
	Status []byte
	Fit    []byte
	Start  []byte
	Size   []byte
	Next   []byte
	Name   []byte
}

type Partition struct {
	Status []byte
	Type   []byte
	Fit    []byte
	Start  []byte
	Size   []byte
	Name   []byte
}

type MountedPartition struct {
	Letra         string
	Identificador string
	State         int
	Type          int
	Name          string
}

type Discos struct {
	Path       string
	Number     int
	State      int
	Partitions [10]MountedPartition
}

type SuperBloque struct {
	FilesystemType  []byte
	InodesCount     []byte
	BlocksCount     []byte
	FreeBlocksCount []byte
	FreeInodesCount []byte
	Mtime           []byte
	MntCount        []byte
	Magic           []byte
	InodeSize       []byte
	BlockSize       []byte
	FirstInode      []byte
	FirstBlock      []byte
	BmInodeStart    []byte
	BmBlockStart    []byte
	InodeStart      []byte
	BlockStart      []byte
}

type Inodo struct {
	UID   []byte
	GID   []byte
	Size  []byte
	Atime []byte
	Ctime []byte
	Mtime []byte
	Block []byte
	Type  []byte
	Perm  []byte
}

func NewInodo(UID, GID, Size, Type, Perm int) *Inodo {
	timeStart := []byte(time.Now().Format(time.RFC850))
	var block [16]int
	for i := range block {
		block[i] = -1
	}
	Block := IntArrayToBytes(block)
	return &Inodo{IntToBytes(UID), IntToBytes(GID), IntToBytes(Size), timeStart, timeStart, timeStart, Block, IntToBytes(Type), IntToBytes(Perm)}
}

type BloqueCarpeta struct {
	BContent [4]Content
}

type Content struct {
	Name  []byte
	Inodo []byte
}

type BloqueArchivo struct {
	Content []byte
}

func ExistName(name string, mbr MBR) bool {
	for i := 0; i < 4; i++ {
		if string(mbr.Partition[i].Name) == name {
			return true
		}
	}
	return false
}

func BytesToInt(bytes []byte) int {
	value, _ := strconv.Atoi(string(bytes))
	return value
}

func IntToBytes(num int) []byte {
	value := []byte(strconv.Itoa(num))
	return value
}

func IntArrayToBytes(intArray [16]int) []byte {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	if err := enc.Encode(intArray); err != nil {
		log.Fatal(err)
	}
	return buff.Bytes()
}

func BytesToIntArray(bytesArray []byte) [16]int {
	var value [16]int
	buff := bytes.NewBuffer(bytesArray)
	dec := gob.NewDecoder(buff)
	if err := dec.Decode(&value); err != nil {
		log.Fatal(err)
	}
	return value
}
