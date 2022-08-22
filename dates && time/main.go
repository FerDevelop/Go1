package main

import (
	"fmt"
	"time"
)

func main() {

	variableTiempo := time.Now()
	fmt.Println("este evento sucede en este momento\n", variableTiempo)

	anioMesDiaHoraMinSegMilUbicacion := time.Date(2022, time.August, 21, 23, 0, 0, 0, time.UTC)
	fmt.Println("\nEsta fecha la invente con formato", anioMesDiaHoraMinSegMilUbicacion)

	fmt.Printf(anioMesDiaHoraMinSegMilUbicacion.Format(time.ANSIC))

	tiempoConvertido, _ := time.Parse(time.ANSIC, "Sun Aug 21 23:00:00 2022")
	fmt.Printf("el tiempo parseado es %T\n", tiempoConvertido)
}
