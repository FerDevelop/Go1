package main

import "fmt"

func main() {
	/* los MAPAS: son estructuras de datos o coleccion de datos, son elementos desordenados
	que se asemejan a los mapas en python, o se equiparan a los objetos literales en javascript
	estos llevan: CLAVE Y VALOR */

	dias := make(map[int]string)

	fmt.Println(dias)

	//agregar datos al map
	dias[1] = "domingo"
	dias[2] = "lunes"
	dias[3] = "martes"
	dias[4] = "miercoles"
	dias[5] = "jueves"
	dias[6] = "viernes"
	dias[7] = "sabado"

	fmt.Println(dias, len(dias))

	dias[7] = "SABADO"
	fmt.Println(dias, len(dias))
	// DELETE es un metodo para borrar elementos de un mapa, el primer argumento es el MAPA y el segundo LA CLAVE
	// la capacidad no se puede mostrar, solo la longitud
	delete(dias, 7)

	fmt.Println(dias, len(dias))

	// se puede crear un mapa el cual su VALOR sea un ARREGLO
	// NUEVO MAPA que va a tener un arreglo como VALOR
	estudiantes := make(map[string][]int)

	fmt.Println(estudiantes, len(estudiantes))

	estudiantes["fernando"] = []int{38, 39, 49}
	estudiantes["medidas"] = []int{183, 86, 2023}
	fmt.Println(estudiantes, len(estudiantes))
	fmt.Println(estudiantes["fernando"])
	fmt.Println(estudiantes["fernando"][1]) // ahora quiero acceder al indice 1 de este arreglo

}
