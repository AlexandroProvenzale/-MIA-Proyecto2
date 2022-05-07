package Program

import "fmt"

func PausarPrograma() {
	var tecla string
	fmt.Println("----------------------------------------------------------")
	fmt.Println("Pausado")
	fmt.Println("Presiona una tecla")
	fmt.Println("----------------------------------------------------------")
	fmt.Scanf("%\n", &tecla)
	return
}
