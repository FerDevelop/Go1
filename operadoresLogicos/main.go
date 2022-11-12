package main

import "fmt"

func main() {
	//Not
	//fmt.Println(!false)
	//And - ambas deben ser verdaderas
	//fmt.Println(true && false)
	//Or -  al menos una de ellas es verdadera
	//mayor o igual que
	//fmt.Println(true || false)

	b := 2

	r := b == 2 || !(4 > b)
	fmt.Println(r)
}
