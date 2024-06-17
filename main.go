package main

import (
	e "github.com/KronosSolucionesTIC/godesde0/ejer_interfaces"
	"github.com/KronosSolucionesTIC/godesde0/modelos"
)

func main() {
	Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)
}
