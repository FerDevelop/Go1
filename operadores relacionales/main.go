package main

import "fmt"

func main() {
	// estructuras de control -  control de flujos.
	//operadores relacionales o de comparacion.

	a := 4
	b := 5

	var r bool
	//igualdad
	r = a == b
	fmt.Printf("%v es igual que %d? %t \n ", a, b, r)

	//distinto
	r = a != b
	fmt.Printf("%v es distinto que %d? %t \n ", a, b, r)

	//mayor que
	r = a > b
	fmt.Printf("%v es mayor que %d? %t \n ", a, b, r)

	//menor que
	r = a < b
	fmt.Printf("%v es menor que %d? %t \n ", a, b, r)

	//mayor o igual que
	r = a >= b
	fmt.Printf("%v es mayor o igual que %d? %t \n ", a, b, r)

	//menor o igual que
	r = a <= b
	fmt.Printf("%v es menor o igual que %d? %t \n ", a, b, r)

}
