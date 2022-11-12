package main

import "fmt"

func main() {
	// FOR EACH utilizado para recorrer colecciones de datos (ARRAY-SLICE-MAP)

	//FOR normal
	nombres := [...]string{"julio", "fernando", "lepore"}
	/*        for iterador := 0; iterador < len(nombres); iterador++ {
		fmt.Println(nombres[iterador])
		//con esto va a devolver elemento por elemento listados
	}   */

	//FOR EACH
	/* esto se equipara al for tradicional con el iterador ya que
	es un metodo con 3 argumentos (indice, elemento y el arreglo que recorre
		y devuelto por el metodo range) */
	for indice, elemento := range nombres {
		fmt.Println(indice, elemento)
	}

	/* aca al usar guion bajo, deja sin uso el primer argumento,
	que seria el INDICE retornado del metodo range */
	for _, elementoDeArray := range nombres {
		fmt.Println(elementoDeArray)
	}
	/* aca no se detalla el segundo argumento del metodo range que seria
	el elemento que se recorre y solo se pasa el primer argumento en el print */
	for indiceDelArreglo := range nombres {
		fmt.Println(indiceDelArreglo)
	}
}
