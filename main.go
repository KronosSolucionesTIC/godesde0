package main

import (
	"fmt"

	"github.com/KronosSolucionesTIC/godesde0/variables/ejercicios"
)

func main() {
	numero, texto := ejercicios.ConvNumerico("a")
	fmt.Println(numero)
	fmt.Println(texto)
}
