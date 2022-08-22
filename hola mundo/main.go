package main

import "fmt"

const miAltura int = 183

var saludo string = "hola"
var nombre string = "fernando"

func main() {
	fmt.Println(saludo)
	fmt.Println(nombre)
	fmt.Printf("el tipo de variable es: %T %T\n", saludo, nombre)

	var numEntero int = 40
	fmt.Println("el valor de la variable es:", numEntero)
	var valorPorDefecto int
	fmt.Println("el valor de la variable por defecto es:", valorPorDefecto)

	miCadena := "esta es otra forma de declarar variables con :=\n"
	fmt.Println(miCadena)
	fmt.Printf("funcion Printf: la variable es del tipo %T", miCadena)
	fmt.Println("\nMi estatura es:", miAltura, "cm y no va a cambiar")
	fmt.Printf("el tipo de la variable miAltura es %T", miAltura)
}
