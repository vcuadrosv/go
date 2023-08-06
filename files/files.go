package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/vcuadrosv/go/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.OtroNumero()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear un archivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumTabla() {
	var texto string = ejercicios.OtroNumero()
	if !Append(fileName, texto) {
		fmt.Println("Error al conectar el contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	//WRONLY Son los parametros de escritura y O_APPEND son de guardado, con el chmod 0644 permisos
	// que se dan para abrir cualquier archivo
	if err != nil {
		fmt.Println("Error durante el append" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto) //perismos para grabar en el archivo
	if err != nil {
		fmt.Println("Error durante el WriteString" + err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error durante la lectura" + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println(">" + registro)
	}
	archivo.Close()

}
