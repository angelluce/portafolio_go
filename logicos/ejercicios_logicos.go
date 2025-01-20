package logicos

import "fmt"

func Menu_logicos() {
	var opcion int
	fmt.Println("** Submenú: Ejercicios lógicos **")
	fmt.Println("1. Verificar vocal")
	fmt.Println("2. Números pares e impares")
	fmt.Println("3. Factorial")
	fmt.Println("4. Sumar números hasta llegar a 50")
	fmt.Println("5. Calcular área")
	fmt.Println("Ingrese el número del ejercicio:")
	fmt.Scan(&opcion)

	switch opcion {
	case 1:
		veficar_vocal()
	case 2:
		numeros_pares_impares()
	case 3:
		factorial()
	case 4:
		suma_hasta_50()
	case 5:
		calcular_area()
	default:
		fmt.Println("Opción no válida")
	}
}

// Crear un aplicativo que pueda adivinar si es Vocal o No. Para eso una letra tiene que ser ingresada por teclado
func veficar_vocal() {
	fmt.Println("8. Verificar vocal")
	var letra string
	fmt.Println("Ingrese una letra: ")
	fmt.Scan(&letra)
	switch letra {
	case "a", "e", "i", "o", "u":
		println("Es una vocal minúscula")
	case "A", "E", "I", "O", "U":
		println("Es una vocal mayúscula")
	default:
		println("No es una vocal")
	}
}

// Solicita un número e imprime todos los números pares e impares desde 1 hasta ese número con el mensaje "es par"
// o "es impar" si el número es 5 el resultado será: 1 - es impar 2 - es par 3 - es impar 4 - es par 5 - es impar
func numeros_pares_impares() {
	fmt.Println("9. Números pares e impares")
	var n int
	fmt.Println("Ingrese un número: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			println(i, "es par")
		} else {
			println(i, "es impar")
		}
	}
}

// Escriba un programa que pida un número entero mayor que cero y calcule su factorial. La factorial es el resultado de
// multiplicar ese número por sus anteriores hasta la unidad
func factorial() {
	fmt.Println("10. Factorial")
	var n, resultado int
	fmt.Println("Ingrese un número: ")
	fmt.Scan(&n)
	resultado = 1
	for i := 1; i <= n; i++ {
		resultado *= i
		fmt.Println("Iteración", i, ":", resultado)
	}
	fmt.Println("El factorial de", n, "es", resultado)
}

// Escribe un programa que permita ir introduciendo una serie indeterminada de números mientras su suma no supere 50.
// Cuando esto ocurra, se debe mostrar el total acumulado y el contador de cuantos números se han introducido
func suma_hasta_50() {
	fmt.Println("11. Sumar números hasta llegar a 50")
	var n, suma, contador int
	suma = 0
	contador = 0
	for suma <= 50 {
		fmt.Println("Ingrese un número: ")
		fmt.Scan(&n)
		suma += n
		contador++
	}
	fmt.Println("Total acumulado:", suma)
	fmt.Println("Cantidad de números ingresados:", contador)
}

// Escribir un programa con un bucle infinito con opciones para elegir, que pueda calcular el área de 2 figuras geométricas,
// triángulo y rectángulo. En primer lugar, pregunta de qué figura se quiere calcular el área, después solicita los datos que
// necesites para calcularlo. 1) triángulo = b * h/2 2) rectángulo = b * h 3) Salir: Solo cuando escojas esta opción el bucle se detendrá
func calcular_area() {
	fmt.Println("12. Calcular área")
	var option int

	fmt.Println("Ejercicio\nCalcular el área de un triángulo o un rectángulo")
	fmt.Println("1) Triángulo")
	fmt.Println("2) Rectángulo")
	fmt.Println("3) Salir")
	fmt.Println("Ingrese una opción: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		var base, altura float64
		fmt.Println("Ingrese la base: ")
		fmt.Scan(&base)
		fmt.Println("Ingrese la altura: ")
		fmt.Scan(&altura)
		area := base * altura / 2
		fmt.Println("El área del triángulo es", area)

	case 2:
		var base, altura float64
		fmt.Println("Ingrese la base: ")
		fmt.Scan(&base)
		fmt.Println("Ingrese la altura: ")
		fmt.Scan(&altura)
		area := base * altura
		fmt.Println("El área del rectángulo es", area)

	case 3:
		fmt.Println("Saliendo...")

	default:
		fmt.Println("Opción inválida")
	}
}
