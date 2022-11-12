package main

import "fmt"

func main() {
	//ahora voy a crear un slice vacio desde cero
	// make es un metodo o una funcion ya definida en Go

	numeros := make([]int, 3, 3)

	// primer argunmento es el slice vacio que va a contener eneteros
	// segundo argumento es la longitud que tendra el slice
	//el tercer argumento es la capacidad que tendra el slice

	fmt.Println(numeros, len(numeros), cap(numeros))

	numeros[0] = 100
	numeros[1] = 200
	numeros[2] = 300

	numeros = append(numeros, 400)

	fmt.Println(numeros, len(numeros), cap(numeros))
}
