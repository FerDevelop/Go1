package main

import (
	"fmt"
	"strings"
)

func revertir(cadena string) string {
	arrayCadena := strings.Split(cadena, "") // convierto una cadena en slice(array)
	arrayInvertida := make([]string, 0)      // crea un slice
	for iterador := len(arrayCadena) - 1; iterador >= 0; iterador-- {
		arrayInvertida = append(arrayInvertida, arrayCadena[iterador])
	} // este iterador me va llenando el slice con elementos del arrayCadena
	return strings.Join(arrayInvertida, "") // convierte en string el arreglo
}
func esCapicua(palabra string) bool {

	palabra = strings.ToLower(palabra)             // convierte el string en minuscula
	palabra = strings.ReplaceAll(palabra, " ", "") // reemplaza o remueve espacios

	palabraInvertida := revertir(palabra)
	return palabra == palabraInvertida // esta funcion va a retornar un BOOLEANO
}

func main() {

	if esCapicua("DANIELA") {
		fmt.Println("la palabra introducida es CAPICUA")
	} else {
		fmt.Println("la palabra introducida no es CAPICUA")
	} // o sea si la funcion esCapicua da true en

	/* llamo a la funcion esCapicua la cual ya trabaja
	por dentro la funcion revertir pasada como metodo en el print de la funcion
	esCapicua */

}
