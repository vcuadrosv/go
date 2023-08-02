package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var Numero1 int
var Numero2 int
var leyenda string

func IngresoNumeros() {
	var err error // Declaración de la variable err en el alcance de la función
	fmt.Println("Ingrese numero 1: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		Numero1, err = strconv.Atoi(scanner.Text()) // Asignación del valor convertido a numero1
		if err != nil {
			panic("El dato ingresado es incorrecto: " + err.Error())
		}
	}

	fmt.Println("Ingrese numero 2: ")
	if scanner.Scan() {
		Numero2, err = strconv.Atoi(scanner.Text()) // Asignación del valor convertido a numero2
		if err != nil {
			panic("El dato ingresado es incorrecto: " + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda: ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	resultado := Numero1 * Numero2
	fmt.Println(leyenda, resultado)
}
