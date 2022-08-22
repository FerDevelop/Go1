package main

import "fmt"

var saludo string = "hola"
var nombre string = "fernando"

func main() {
	fmt.Println(saludo)
	fmt.Println(nombre)
	fmt.Printf("el tipo de variable es: %T %T", saludo, nombre)
}
