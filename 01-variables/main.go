//package main es el paquete principal de go y es el que se ejecuta primero.
//siempre debe de ir en el primer renglon
package main

import "fmt"

//Test de la funcion main
var Test = 123
// func main() {
	// 	name, age := "Ivan", 20
	// 	fmt.Println("Hola", name, ". Tienes", age, "años.")
	// }
	
func main() {
	// PI es una constante que vale 3.1416
	const pi = 3.1416

	// diametro es una variable que vale 10
	var diametro = 10
	var area = pi * float64(diametro)
	fmt.Println("Area:", area)
}