package gorutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLentoo(nombre string) {
	letras := strings.Split(nombre, " ")
	for _, letra := range letras {
		time.Sleep(10000 * time.Millisecond)
		fmt.Println(letra)
	}

}
