package main
import "fmt"

func main() {
	fmt.Println("Slices")
	// Declaración de un slice
	// set := [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	// animales := set[0:3]

	// personas := []string{"Juan", "Pedro", "Luis", "Ana", "Maria", "Jose"}
	// pares := personas[0:6]
	// pares[0] = "Juan Carlos"
	// ultimos := personas[4:6]

	// fmt.Println(animales) // [a b c]
	// fmt.Println(pares[:]) // Slice de todos los elementos
	// fmt.Println(ultimos)  // [Maria Jose]

	comidas := []string{"Pizza", "Hamburguesa", "Tacos", "Enchiladas", "Sushi"}
	rapida := comidas[0:2]
	fmt.Println(len(rapida)) // [Pizza Hamburguesa]

}