package main

import "fmt"

func main() {
	// palabras reservadas que se usan dentro de los BUCLES (CONTINUE y BREAK)
	// CONTINUE hace saltar la iteracion del bucle
	// BREAK rompe con el bucle o termina la ejecucion del bucle
	for variable := 0; variable <= 10; variable++ {
		if variable == 5 {
			fmt.Println("salta la interacion Y NO SE VE LA ITERACION 5")
			continue
		}
		if variable == 8 {
			fmt.Println("romper con bucle con esta variable ", variable)
			break // la palabra reserveda va al final
		}
		fmt.Println(variable)
	}

}
