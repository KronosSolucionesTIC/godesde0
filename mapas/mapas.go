package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)

	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"
	paises["Colombia"] = "Bogota"

	fmt.Println(paises)
	fmt.Println(paises["Colombia"])

	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 40,
		"Liverpool":   50,
	}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Println(equipo, puntaje)
	}

	delete(campeonato, "Barcelona")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Liverpool"]
	fmt.Printf("El puntaje %d, y el equipo existe %t \n", puntaje, existe)
}
