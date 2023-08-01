package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.888
	Fecha = time.Now().Local()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConvirtiendoaTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
