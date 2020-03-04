package main

import "fmt"

var numero int
var texto string
var status bool = true

func main() {
	var numero2, numero3, numero4, texto, status = 2, 3, 4, "daniel", false

	var moneda float64 = 0
	numero2 = int(moneda)

	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
	fmt.Println(status)

	mostrarStatus()
}
func mostrarStatus() {
	fmt.Println(status)
}
