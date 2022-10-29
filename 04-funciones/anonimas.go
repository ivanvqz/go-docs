package main
import "fmt"

func main()  {
	// Funciones anónimas
	// saludo := func() {
	// 	fmt.Println("Hola mundo")
	// }
	// saludo()

	//funciones anonimas autoejecutadas
	func(nombre string) {
		fmt.Println("Hola " + nombre)
	}("Ivan")
}