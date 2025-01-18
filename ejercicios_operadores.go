package main

import "fmt"

func ejercicio_aritmeticos() {
	fmt.Println("Operaciones aritméticas")
	var a, b, resultado int

	// Lectura de datos
	fmt.Println("Ingrese el valor de a: ")
	fmt.Scan(&a)

	fmt.Println("Ingrese el valor de b: ")
	fmt.Scan(&b)

	// Operaciones aritméticas
	resultado = a + b
	println("a + b =", resultado)
	resultado = a - b
	println("a - b =", resultado)
	resultado = a * b
	println("a * b =", resultado)
	resultado = a / b
	println("a / b =", resultado)
	resultado = a % b
	println("a % b =", resultado)
}

func ejercicio_relacionales() {
	fmt.Println("Operaciones relacionales")
	var a, b int

	// Lectura de datos
	fmt.Println("Ingrese el valor de a: ")
	fmt.Scan(&a)

	fmt.Println("Ingrese el valor de b: ")
	fmt.Scan(&b)

	// Operaciones relacionales
	println("a == b", a == b)
	println("a != b", a != b)
	println("a > b", a > b)
	println("a < b", a < b)
	println("a >= b", a >= b)
	println("a <= b", a <= b)
}

func ejercicio_asignacion() {
	fmt.Println("Operaciones de asignación")
	var a, b, resultado int

	// Lectura de datos
	fmt.Println("Ingrese el valor de a: ")
	fmt.Scan(&a)

	fmt.Println("Ingrese el valor de b: ")
	fmt.Scan(&b)

	// Operaciones de asignación
	resultado = a
	println("a = a", resultado)
	resultado += b
	println("a += b", resultado)
	resultado -= b
	println("a -= b", resultado)
	resultado *= b
	println("a *= b", resultado)
	resultado /= b
	println("a /= b", resultado)
	resultado %= b
	println("a %= b", resultado)
}

func ejercicio_condicionales() {
	fmt.Println("Operaciones condicionales")
	var a, b int

	// Lectura de datos
	fmt.Println("Ingrese el valor de a: ")
	fmt.Scan(&a)

	fmt.Println("Ingrese el valor de b: ")
	fmt.Scan(&b)

	// Operaciones condicionales
	if a == b {
		println("a es igual a b")
	} else if a > b {
		println("a es mayor que b")
	} else if a < b {
		println("a es menor que b")
	} else {
		println("a no es igual a b")
	}
}

func ejercicio_logicos() {
	fmt.Println("Operadores lógicos")
	var a, b int

	// Lectura de datos
	fmt.Println("Ingrese el valor de a: ")
	fmt.Scan(&a)

	fmt.Println("Ingrese el valor de b: ")
	fmt.Scan(&b)

	// Operadores lógicos
	if !(a == b) {
		println("a no es igual a b")
	} else if a == 0 && b == 0 {
		println("a y b son iguales a 0")
	} else if a > 0 || b > 0 {
		println("a no es igual a b")
	} else {
		println("Ninguna de las anteriores")
	}
}

func ejercicio_switch() {
	fmt.Println("Switch")
	var a int

	// Lectura de datos
	fmt.Println("Ingrese el valor de a (1-5): ")
	fmt.Scan(&a)

	// Switch
	switch a {
	case 1:
		println("Uno")
	case 2:
		println("Dos")
	case 3:
		println("Tres")
	case 4:
		println("Cuatro")
	case 5:
		println("Cinco")
	default:
		println("Ninguna de las anteriores")
	}
}

func ejercicio_for() {
	fmt.Println("For")

	var n int
	fmt.Print("Ingresa un número: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		println(i)
	}
}
