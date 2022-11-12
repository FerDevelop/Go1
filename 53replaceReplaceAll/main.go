package main

import (
	"fmt"
	"strings"
)

func capicua(palabra string) {
	fmt.Println(palabra)
	palabra = strings.ToUpper(palabra)
	palabra = strings.ReplaceAll(palabra, " ", "") /* esto se usa para
	donde hay un espacio (segundo argumento) y en el tercer argumento
	se coloca "" (es decir sin espacios) */

	/* palabra = strings.Replace(palabra, "S", "Z", 4)  "z" es el caracter que esta
	escrito	"s" es el caracter que va a reemplazar el an terior y 4 es la cantidad
	de veces que reemplazara el caracter que meto en lugar*/
	fmt.Println(palabra)
}

func main() {
	/*ahora a trabajar con palabras palindroma o capicua, es decir que se lee
	de igual forma o al reves, creando una funcion de string */
	capicua("La Lus Asul y el sinc, asi como el sumo...")
}
