package Program

import "fmt"

func PausarPrograma() {
	var tecla string
	fmt.Println("----------------------------------------------------------")
	fmt.Println("Pausado")
	fmt.Println("Presiona una tecla")
	fmt.Println("----------------------------------------------------------")
	_, _ = fmt.Scanf("%s\n", &tecla)
	return
}
