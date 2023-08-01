package main

import (
	"fmt"

	"github.com/vcuadrosv/go/variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConvirtiendoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
