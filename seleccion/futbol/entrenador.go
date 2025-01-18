package futbol

import "fmt"

type Entrenador struct {
	Seleccion
	Federacion string
}

func (e Entrenador) DirigirPartido() {
	fmt.Printf("El entrenador %s dirige el partido\n", e.Nombre)
}
