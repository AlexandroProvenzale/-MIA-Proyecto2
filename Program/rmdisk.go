package Program

import (
	"fmt"
	"os"
	"strings"
)

func RemoveDisk(path string) {
	fmt.Println("-------------------- Eliminar disco --------------------")
	fmt.Println(path)
	if _, err := os.Stat(path); err == nil {
		fmt.Println("¿Seguro que desea borrar el disco en la locación:", path)
		fmt.Println("Ingrese respuesta")
		fmt.Println("S = SI, N = NO")
		for {
			var tecla string
			_, _ = fmt.Scanf("%s\n", &tecla)
			if strings.ToLower(tecla) == "s" {
				if err := os.Remove(path); err != nil {
					fmt.Println("Error al eliminar el archivo", path)
				} else {
					fmt.Println("Disco", path, "eliminado exitosamente!")
					return
				}
			} else if strings.ToLower(tecla) == "n" {
				fmt.Println("No se eliminó el archivo")
				return
			} else {
				fmt.Println("Ingrese una opción correcta")
			}
		}
	} else {
		fmt.Println("Error: No hay ningún disco en la ruta ingresada")
		return
	}
}
