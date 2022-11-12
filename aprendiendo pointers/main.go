package main

import "log"

func main() {
	var miString string
	miString = "verde"

	log.Println("my string se establece en", miString)

	changeUsingPointer(&miString)

	log.Println("despues de llamar la funcion mystring va a cambiar a", miString)
}

func changeUsingPointer(s *string) {
	log.Println("s se establece en el area de memoria", s)
	nuevoValor := "rojo"
	*s = nuevoValor

}
