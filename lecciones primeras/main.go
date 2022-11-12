package main

import "fmt"

func main() {
	fmt.Println("hello, world.")

	var whatToSay string
	var i int
	whatToSay = "good bye, cruel world"
	fmt.Println(whatToSay)
	i = 7
	fmt.Println("i set to", i)
	whatWassaid := saySomething()
	fmt.Println("the function returned", whatWassaid)
}

func saySomething() string {
	return "something"
}
