package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expresion string) int {
	valores := strings.Split(expresion, "+")
	var resultado int
	for _, valor := range valores {
		numero, errorVisible := strconv.Atoi(valor) // convierte string a entero

		if errorVisible != nil { // aqui el error es distinto a nil(valido)
			fmt.Println(errorVisible)
			fmt.Println("la operacion solicitada no es correcta")
		} else {
			resultado += numero // este es el valor de nil
		}
	}
	return resultado
}
func main() {
	var expresion string
	var resultado int
	fmt.Println("colocar numeros que se sumaran:")
	fmt.Scanln(&expresion)
	resultado = sumar(expresion)
	fmt.Printf("=> %d\n", resultado)
}
