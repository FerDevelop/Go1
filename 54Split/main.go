package main

import (
	"fmt"
	"strings"
)

func revertir(cadena string) string {
	arrayCadena := strings.Split(cadena, "")
	fmt.Println(arrayCadena)
	return ""
}
func esCapicua(palabra string) {

	fmt.Println(palabra)
	palabra = strings.ToLower(palabra)
	palabra = strings.ReplaceAll(palabra, " ", "")

	fmt.Println(palabra)
}

func main() {
	esCapicua("Luz azul")
	revertir("luz azul")
}
