package Program

import (
	"fmt"
	"os"
)

func RemoveDisk(path string) {
	fmt.Println("-------------------- Eliminar disco --------------------")
	fmt.Println(path)
	if _, err := os.Stat(path); err == nil {
		err := os.Remove(path)
		if err != nil {
			fmt.Println("Error al eliminar el archivo", path)
		} else {
			fmt.Println("Archivo eliminado exitosamente!")
		}
	} else {
		fmt.Println("Error: No hay ningún disco en la ruta ingresada")
		return
	}
}
