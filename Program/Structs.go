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
	Number        int
	Identificador string
	State         int
	Type          int
}

type Discos struct {
	Path             string
	Letra            string
	State            int
	MountedPartition [10]Partition
}
