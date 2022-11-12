package main

import "fmt"

func main() {
	a := 10

	// suma en asignacion
	// a = a + 2
	// el resultado es 15
	a += 5
	fmt.Println(a)

	// resta en asignacion
	// a = a - 10
	//el resultado es 5
	a -= 10
	fmt.Println(a)

	//multiplicacion en asignacion
	// a= a *2
	// el resultado es 10
	a *= 2
	fmt.Println(a)

	//division en asignacion
	// a = a /=5
	// el resultado es 2

	a /= 5
	fmt.Println(a)

	//modulo en asignacion
	//a= a % 2
	// el resultado es 0 (cero)
	a %= 2
	fmt.Println(a)

}
