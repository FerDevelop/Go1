package main

import (
	"fmt"
	"math"
)

func main() {
	entero1, entero2, entero3 := 22, 35, 60
	sumaEnteros := entero1 + entero2 + entero3
	fmt.Println("la suma de los enteros es: ", sumaEnteros)

	dec1, dec2, dec3 := 52.6, 30.5, 20.8
	sumaDecimales := dec1 + dec2 + dec3
	fmt.Println("la suma decimales: ", sumaDecimales)
	sumaDecimales = math.Round(sumaDecimales*100) / 100 //forma correcta de recortar digitos
	fmt.Println("la suma decimalescon el redondeo proporcionado por el metodo math.Round es: ", sumaDecimales)

	radiodeCirculo := 15.5
	fmt.Println("el radio de circulo es: ", radiodeCirculo)
	circunferencia := radiodeCirculo * 2 * math.Pi
	fmt.Printf("la circunferencia es de: %2f\n ", circunferencia)
}
