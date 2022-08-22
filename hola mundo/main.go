package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	fmt.Printf("el tipo de la variable miAltura es %T\n", miAltura)

	leer := bufio.NewReader(os.Stdin)
	fmt.Print("introduci el texto por favor: ")
	varIngresousuario, _ := leer.ReadString('\n')
	fmt.Println("has ingresado el texto: ", varIngresousuario)

	leerConsola := bufio.NewReader(os.Stdin)
	fmt.Print("introduci un numero por favor: ")
	inputUsuario, _ := leerConsola.ReadString('\n')
	unDecimal, err := strconv.ParseFloat(strings.TrimSpace(inputUsuario), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("el valor ingresado es", unDecimal)
	}

}
