package ejer_interfaces

import (
	"fmt"

	"github.com/KronosSolucionesTIC/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/una %s, y estoy respirando \n", hu.Sexo())
}
