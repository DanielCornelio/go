package main

import "fmt"

var numero int
var texto string
var status bool

func main() {
	var numero2, numero3, numero4, texto = 2, 3, 4, "daniel"

	var moneda float64 = 0
	numero2 = moneda

	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
}
