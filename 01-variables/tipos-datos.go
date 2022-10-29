package main

import "fmt"

func main() {
	//Bool, string, int, float64
	var a = true
	var b = "Hola"
	var c  	uint64 = 12312
	fmt.Printf("Tipo: %T, Valor: %v", a, a)
	//Verbos para imprimir tipo de dato: %T
	//Verbos para imprimir valor: %v
	fmt.Printf("\nTipo: %T, Valor: %v", b, b)
	fmt.Printf("\nTipo: %T, Valor: %v", c, c)
}