package main
import "fmt"

func main()  {
	// saludo("Juan", "Perez")
	estado := "depierto"
	cambiar(&estado)
	fmt.Println(estado)
}

func cambiar(estado *string)  {
	*estado = "dormido"
}

// func saludo(nombre string, apellido string)  {
// 	fmt.Printf("Hola %s %s", nombre, apellido)
// 	return 
// }