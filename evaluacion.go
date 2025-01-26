package main

import "fmt"

func main() {

	const IVA = 0.16

	var name string
	var last_name string
	var cedula int16
	var compra float64

	fmt.Println("Ingrese su nombre:")
	fmt.Scanln(&name)
	fmt.Println("Ingrese el apellido:")
	fmt.Scanln(&last_name)
	fmt.Println("Ingrese su cédula:")
	fmt.Scanln(&cedula)
	fmt.Println("Ingrese la compra total:")
	fmt.Scanln(&compra)

	var precio_total float64 = compra + (compra * IVA)

	fmt.Printf("Nombre completo del cliente: %s %s\n", name, last_name)
	fmt.Printf("Cédula del cliente: %d\n", cedula)
	fmt.Printf("Precio Total de la compra: %.2f", precio_total)

}
