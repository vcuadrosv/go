package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func OtroNumero() int {
	for {
		fmt.Println("Ingrese un número:")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			numero, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("El dato ingresado es incorrecto. Inténtelo nuevamente.")
			} else {
				return numero
			}
		}
	}
}
