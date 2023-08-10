package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 4, 5}
var arreglo [10]int = [10]int{10, 5, 6, 5, 7, 8, 9, 2, 34, 6}

func MuestroSlices() {
	fmt.Println(tablaS)

	porcion := arreglo[3:]   //Slice creado desde un vector, desde la posicion 3
	porcion2 := arreglo[:5]  //Slice creado desde un vector, desde la posicion 0 a 5
	porcion3 := arreglo[6:7] //Slice creado desde un vector, desde la posicion 6 a 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d\n", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d\n", len(nums), cap(nums))
}
