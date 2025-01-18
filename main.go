package main

import "fmt"

func main() {
	var opcion int

	fmt.Println("Portafolio de ejercicios con GO | Angel Lucero")

	fmt.Println("\tMenú de ejercicios")

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

	fmt.Println("Ingrese el número del ejercicio que desea ejecutar: ")
	fmt.Scan(&opcion)

	switch opcion {
	case 1:
		ejercicio_aritmeticos()
	case 2:
		ejercicio_relacionales()
	case 3:
		ejercicio_asignacion()
	case 4:
		ejercicio_logicos()
	case 5:
		ejercicio_condicionales()
	case 6:
		ejercicio_switch()
	case 7:
		ejercicio_for()
	case 8:
		veficarVocal()
	case 9:
		numerosParesImpares()
	case 10:
		factorial()
	case 11:
		sumaHasta50()
	case 12:
		calcularArea()
	default:
		fmt.Println("Opción no válida")
	}
}
