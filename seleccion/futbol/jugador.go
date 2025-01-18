package futbol

import "fmt"

type Jugador struct {
	Seleccion
	Dorsal int
}

func (j Jugador) JugarPartido() {
	fmt.Printf("El jugador %s juega el partido\n", j.Nombre)
}
