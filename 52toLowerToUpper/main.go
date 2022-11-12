package main

import (
	"fmt"
	"strings"
)

func capicua(palabra string) {
	fmt.Println(palabra)
	palabra = strings.ToLower(palabra)
	fmt.Println(palabra)
}

func main() {
	/*ahora a trabajar con palabras palindroma o capicua, es decir que se lee
	de igual forma o al reves, creando una funcion de string */
	capicua("Menem")
}
