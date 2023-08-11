package middelwares

import "fmt"

func suma(a, b int) int {
	return a + b
}

func resta(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func Mimiddelware() {
	fmt.Println("Inicio")

	result := operacionesMidd(suma)(2, 3)
	fmt.Println(result)

	result = operacionesMidd(resta)(6, 3)
	fmt.Println(result)

	result = operacionesMidd(multiplicar)(2, 3)
	fmt.Println(result)
}

func operacionesMidd(f func(int, int) int) func(int, int) int {

	//Codigo del middelware que se ejecuta antes de llamar a la funcion
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)
	}
}
