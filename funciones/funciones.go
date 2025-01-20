package funciones

import (
	"fmt"
)

func Menu_funciones() {
	var opcion int
	fmt.Println("** Submenú: Ejercicios con funciones **")

	fmt.Println("Ingrese el número del ejercicio:")
	fmt.Println("1. Ejercicio con datos")
	fmt.Println("2. Ejercicio saludar")
	fmt.Println("3. Ejercicio semana")
	fmt.Scan(&opcion)

	switch opcion {
	case 1:
		ejercicio_datos()
	case 2:
		ejercicio_saludar()
	case 3:
		ejercicio_semana()
	default:
		fmt.Println("Opción no válida")
	}
}

// Crear 1 arrays números y dos Slicen. El primero tendrá 5 números y el segundo se
// llamará pares y el tercero impares, ambos estarán vacíos. Después multiplica cada
// uno de los números del primer array por un número aleatorio entre 1 y 10, si el
// resultado es par guarda ese número en el Slicen de pares y si es impar en el
// Slicen de impares. Muestra por consola: -la multiplicación que se produce junto
// con su resultado con el formato 2 x 3 = 6 -el Slicen de pares e impares.
func ejercicio_datos() {
	fmt.Println("1. Ejercicio con datos")
	numeros := [5]int{1, 2, 3, 4, 5}
	pares := []int{}
	impares := []int{}
	m := 0

	// Lectura de datos
	for i := 0; i < 5; i++ {
		var aleatorio int
		fmt.Println("Ingrese el valor del número:")
		fmt.Scan(&aleatorio)
		m = numeros[i] * aleatorio
		fmt.Println(numeros[i], "x", aleatorio, "=", m)

		if m%2 == 0 {
			pares = append(pares, m)
		} else {
			impares = append(impares, m)
		}
	}

	fmt.Println("Pares:", pares)
	fmt.Println("Impares:", impares)
}

func ejercicio_saludar() {
	fmt.Println("2. Ejercicio saludar")
	var nombre string
	fmt.Println("Ingrese su nombre: ")
	fmt.Scan(&nombre)
	fmt.Printf("Hola %s, saludos desde la función saludar\n", nombre)
}

func ejercicio_semana() {
	fmt.Println("3. Ejercicio semana")
	semana := map[int]string{
		1: "Domingo",
		2: "Lunes",
		3: "Jueves",
		4: "Miércoles",
		5: "Martes",
		6: "Viernes",
		7: "Sábado",
	}

	for i := 1; i <= len(semana); i++ {
		fmt.Println(i, semana[i])
	}

	var (
		c int
		k int
		v string
	)
	fmt.Println("Corregir los días de la semana")
	fmt.Println("¿Cuantos días quiere corregir?")
	fmt.Scan(&c)

	for i := 0; i < c; i++ {
		fmt.Println("Ingrese la clave:")
		fmt.Scan(&k)
		fmt.Println("Ingrese el valor:")
		fmt.Scan(&v)
		semana[k] = v
	}

	for i := 1; i <= len(semana); i++ {
		fmt.Println(i, semana[i])
	}
}
