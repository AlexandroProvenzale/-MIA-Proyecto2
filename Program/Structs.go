package Program

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

func ExistName(name string, mbr MBR) bool {
	for i := 0; i < 4; i++ {
		if string(mbr.Partition[i].Name) == name {
			return true
		}
	}
	return false
}
