package main

import (
	"fmt"
	funciones "portafolio_go/funciones"
	logicos "portafolio_go/logicos"
	operadores "portafolio_go/operadores"
	seleccion "portafolio_go/seleccion"
)

func main() {
	var opcion int

	fmt.Println("----------------------------------------------")
	fmt.Println("Portafolio de ejercicios con GO | Angel Lucero")

	for {
		fmt.Println("----------------------------------------------")
		fmt.Println("\t** Menú de ejercicios **")

		fmt.Println("1. Ejercicios con operadores")
		fmt.Println("2. Ejercicios lógicos")
		fmt.Println("3. Ejercicios con funciones")
		fmt.Println("4. Ejercicio selección de futbol")

		fmt.Println("0. Salir")
		fmt.Println("----------------------------------------------")

		fmt.Println("Ingrese el número del ejercicio:")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			operadores.Menu_ejercicios()
		case 2:
			logicos.Menu_logicos()
		case 3:
			funciones.Menu_funciones()
		case 4:
			seleccion.Ejercicio_seleccion()
		case 0:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción no válida")
			break
		}
	}
}
