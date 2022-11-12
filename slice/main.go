package main

import "fmt"

func main() {
	//slice
	numeros := []int{1, 2, 3}
	fmt.Println(numeros, len(numeros))
	numeros = append(numeros, 4)
	numeros = append(numeros, 5)

	fmt.Println(numeros, len(numeros))

	//sub-slice
	subSlice := numeros[:2]

	numeros[0] = 100

	fmt.Println(subSlice)
	fmt.Println(numeros)
	subSlice[0] = 1
	fmt.Println(subSlice, len(subSlice))
	fmt.Println(numeros, len(numeros))

	//punteros
	//longitud
	//capacidad

	meses := []string{"enero", "febrero", "marzo"}

	fmt.Printf("len: %v, cap: %v,  %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "abril")

	fmt.Printf("len: %v, cap: %v,  %p \n", len(meses), cap(meses), meses)
}
