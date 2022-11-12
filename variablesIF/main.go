package main

import "fmt"

func main() {
	if /* variables*/ nombre, edad := "fernando", 38; /* ; condicion */ nombre == "fernando" {
		fmt.Println("hola ", nombre)
	} else {
		fmt.Printf("Nombre: %s Edad: %d\n ", nombre, edad)
	}
	//obtener valor de un MAP
	users := make(map[string]string) // creacion de un Map simple

	//agregar valores al map(simple: key-value o sea llave-valor)
	users["fernando"] = "juliofernandolepore@gmail.com"
	users["daniela"] = "benavidez.daniela@gmail.com"
	users["julio"] = "juliocesarlepore@gmail.com"

	fmt.Println(users["juan"]) /* aca llamo una llave(key) inexistente
	esto va a devolver un valor vacio, y cuando solicitamos un valor a un MAPA este nos devuelve
	2 valores y almacenarlo en una variable , los 2 valores que devuelve VALUE y un BOOLEANO */

	//   correo, ok := users["juan"]  // aca en estas 2 variables se almacenan el VALOR
	// segun la llave que puse, y valor true o false que se almacenara en la segunda variable

	// fmt.Println(correo, ok) // en caso de no existir la llave, devolvera un FALSE

	if correo, ok := users["fernando"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("no fue posible obtener el valor del MAP porque la llave solicitada no existe")
	}

}
