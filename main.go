package main

import (
	"fmt"
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

		fmt.Println("1. Operadores aritméticos")
		fmt.Println("2. Operadores relacionales")
		fmt.Println("3. Operadores de asignación")
		fmt.Println("4. Operadores lógicos")
		fmt.Println("5. Condicionales")
		fmt.Println("6. Switch")
		fmt.Println("7. For")

		fmt.Println("8. Verificar vocal")
		fmt.Println("9. Números pares e impares")
		fmt.Println("10. Factorial")
		fmt.Println("11. Sumar números hasta llegar a 50")
		fmt.Println("12. Calcular área")

		fmt.Println("13. Ejercicio selección de futbol")

		fmt.Println("0. Salir")
		fmt.Println("----------------------------------------------")

		fmt.Println("Ingrese el número del ejercicio:")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			operadores.Ejercicio_aritmeticos()
		case 2:
			operadores.Ejercicio_relacionales()
		case 3:
			operadores.Ejercicio_asignacion()
		case 4:
			operadores.Ejercicio_logicos()
		case 5:
			operadores.Ejercicio_condicionales()
		case 6:
			operadores.Ejercicio_switch()
		case 7:
			operadores.Ejercicio_for()
		case 8:
			logicos.VeficarVocal()
		case 9:
			logicos.NumerosParesImpares()
		case 10:
			logicos.Factorial()
		case 11:
			logicos.SumaHasta50()
		case 12:
			logicos.CalcularArea()
		case 13:
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
