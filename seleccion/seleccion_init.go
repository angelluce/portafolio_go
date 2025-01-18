package seleccion

import (
	"fmt"
	futbol "portafolio_go/seleccion/futbol"
)

func Ejercicio_seleccion() {
	fmt.Println("13. Ejercicio selección de futbol")

	entrenador := new(futbol.Entrenador)
	entrenador.Nombre = "Sebastian Beccacece"
	entrenador.Edad = 40
	entrenador.Federacion = "ECU"
	fmt.Println(entrenador)
	entrenador.DirigirPartido()

	jugador1 := new(futbol.Jugador)
	jugador1.Nombre = "Piero Hincapie"
	jugador1.Edad = 21
	fmt.Println(jugador1)
	jugador1.JugarPartido()

	jugador2 := new(futbol.Jugador)
	jugador2.Nombre = "Moises Caicedo"
	jugador2.Edad = 20
	fmt.Println(jugador2)
	jugador2.JugarPartido()
}
