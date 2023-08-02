package ejercicios

import (
	"strconv"
)

func DevolverDosValores(parametro string) (int, string) {
	numero, err := strconv.Atoi(parametro)
	if err != nil {
		// Si hubo un error en la conversión, retornamos el valor cero y el mensaje de error
		return 0, "Error en la conversión" + err.Error()
	}
	if numero > 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "Es menor a 100"
	}
}
