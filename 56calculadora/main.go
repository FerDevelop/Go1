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
		numero, _ := strconv.Atoi(valor)

		resultado += numero
	}

	return resultado
}

func main() {
	var expresion string
	var resultado int

	fmt.Println("colocar los numeros que se sumaran =>")
	fmt.Scanln(&expresion)

	sumar(expresion)
	resultado = sumar(expresion)
	fmt.Printf("=> %d\n", resultado)
}
