package main

import (
	"fmt"

	"github.com/vcuadrosv/go/ejercicios"
	/*"runtime"
	"github.com/vcuadrosv/go/variables"
	"github.com/vcuadrosv/go/teclado"*/)

func main() {
	/*variables.MuestroEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConvirtiendoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)*/
	/*if os := runtime.GOOS; os == "linux" || os == "darwin" {
		fmt.Println("Sistema Operativo: ", os)
	} else {
		fmt.Println("Esto es Windows")
	}*/
	/*Para realizarlo con multiples variables*/
	/*switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es:", os)
	case "darwin":
		fmt.Println("Esto es:", os)
	default:
		fmt.Printf("%s \n", os)*/

	/*Ejercicio

	numero, mensaje := ejercicios.DevolverDosValores("dsf")
	fmt.Println("Número:", numero)
	fmt.Println("Mensaje:", mensaje)

	teclado.IngresoNumeros()

	// La multiplicación ahora está dentro de la función IngresoNumeros
	resultado := teclado.Numero1 * teclado.Numero2
	fmt.Println("Resultado:", resultado)*/

	/*for {
		fmt.Println("hola")
		break
	}

	iteraciones.Iterar()*/

	numero := ejercicios.OtroNumero()
	fmt.Println("Tabla del", numero)
	for i := 1; i <= 10; i++ {
		resultado := numero * i
		fmt.Printf("%d x %d = %d\n", numero, i, resultado)
	}

}
