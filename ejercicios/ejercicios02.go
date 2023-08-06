package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func OtroNumero() string {
	scanner := bufio.NewScanner(os.Stdin)
	var texto string

	for {
		fmt.Println("Ingrese un numero: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, numero*i)
	}

	return texto

}

/* se hara una modificacion  ya que e intentara importar el archivo aqui expuesto en
un documento de texto

codigo original

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

*/
