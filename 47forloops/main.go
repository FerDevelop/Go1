package main

import "fmt"

func main() {
	//bucles usando la palabra reservada "for"
	/* -----------------bucle infinito-----------------

	for {
		fmt.Println("hola mundo")
	} */

	/*    -----------------bucle tipo while-----------------------
	numeros := 12455967
	contador := 0

	for numeros > 0 {
		numeros /= 10
		contador++ //se va a ir incrementando
	}
	fmt.Println("la cantidad de digitos es", contador)

	*/

	//BUCLE FOR
	for variable := 0; variable <= 100; variable++ {
		if variable%2 == 0 {
			fmt.Println(variable)
		}
	}
}
