package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{10, 5, 6}
var Matriz [5][5]int

func MuestroArreglo() {
	tabla[7] = 2
	tabla[2] = 7

	tabla2 := [10]string{"Pablo", "Hola"}

	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	Matriz[2][2] = 15
	fmt.Println(Matriz)
}
