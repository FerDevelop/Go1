package main

import "log"

var s = "siete"

func main() {
	var s2 = "seis"

	log.Println("s es", s)
	log.Println("s2 es", s2)

	decirAlgo("xxx")

}

func decirAlgo(s3 string) (string, string) {
	log.Println("s de la funcion decir algo es", s)
	return s3, "world"

}
