package main

import "fmt"

func main() {
	fruta := "manzana"
	// Almacena la dirección en memoria de un string
	var p *string = &fruta
	fmt.Printf("Tipo: %T, Valor: %s, Dirección: %v", fruta, fruta, &fruta)
	fmt.Printf("Tipo: %T, Valor: %v, Desferenciación: %s\n",p, p, *p)
}