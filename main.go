package main

import (
	"fmt"

	a "github.com/vcuadrosv/go/gorutines"
	/*"github.com/vcuadrosv/go/new"*/ /*"runtime"
	"github.com/vcuadrosv/go/variables"
	"github.com/vcuadrosv/go/teclado"
	*/)

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

	//println(ejercicios.OtroNumero())

	//files.SumTabla()

	//files.LeoArchivo()

	//arreglos_slices.Capacidad()

	//mapas.MostrarMapas()

	//users.AltaUsuarios()

	//pedro := new(modelos.Hombre)
	//e.HumanosRespirando(pedro)

	//andrea := new(modelos.Mujer)
	//e.HumanosRespirando(andrea)

	//h.EjemploPanic()

	go a.MiNombreLentoo("valery Cuadros")
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)

}
