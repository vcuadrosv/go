package funciones

import "fmt"

func Calculos() {

	var numero3 int = 31
	var numero4 int = 40

	contador := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}
	fmt.Println(contador(10, 25))

	contador = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero3
	}
	fmt.Println(contador(10, 25))

}
