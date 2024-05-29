package arreglos_slices

import "fmt"

var tablaSlice []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 1, 2, 3, 4, 5, 9}

func MuestroSlice() {
	fmt.Println(tablaSlice)

	porcion := arreglo[3:]   // Toma desde la posicion 3 hasta el final
	porcion2 := arreglo[:5]  // Toma desde la posicion 0 hasta la 5
	porcion3 := arreglo[4:7] // Toma desde la posicion 4 hasta la 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))
}
