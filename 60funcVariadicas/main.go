package main

import "fmt"

/* funciones variadicas son aquellas que reciben numero indeterminados de
argumentos a lo cual se tienen que colocar 3 puntitos (...) */

func sumar(numeros ...int) int {
	// esto la va a recibir como un arreglo
	var total int
	for _, numeros := range numeros {
		total += numeros
	}
	return total

}

func main() {
	resultadoFinal := sumar(1, 2, 3, 4)
	fmt.Println(resultadoFinal)
}
