package main

import "fmt"

func main() {
	/* app para restaurante
	100 dolares de consumo - 10% de descuento
	200 dolares de consumo - 20% de descuento
	mas de 200 dolares de consumo - 30% de descuento
	IGV del 19% (impuesto general de ventas de peru)
	*/
	var consumo, descuento float64
	var datosDescuento string
	// enrada de datos
	fmt.Print("ingrese importe del consumo efectuado: ")
	fmt.Scanln(&consumo)

	if consumo >= 0 {
		// if anidado
		if consumo <= 100 {
			//descuento del 10%
			datosDescuento = "10%"
			descuento = consumo * 0.10
		} else if consumo > 100 && consumo <= 200 {
			//descuento del 20%
			datosDescuento = "20%"
			descuento = consumo * 0.20

		} else if consumo > 200 {
			//descuento del 30%
			datosDescuento = "30%"
			descuento = consumo * 0.30

		}

	} else {
		fmt.Println("error al ingresar el consumo")
	}

	//operaciones
	//monto de descuento
	montoDescuento := consumo - descuento
	igv := montoDescuento * 0.19
	totalPagar := montoDescuento + igv

	//salida de datos
	fmt.Println("-------------factura de consumo-------------------")
	fmt.Println(" descuentos que aplica", datosDescuento)
	fmt.Println("el consumo es", consumo)
	fmt.Println("monto con el descuento realizado:", montoDescuento)
	fmt.Println("impuesto general de ventas", igv)
	fmt.Println("el total a pagar es:", totalPagar)
}
