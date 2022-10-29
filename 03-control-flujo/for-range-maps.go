package main
import "fmt"

func main()  {
	// deportes := map[string]string{"futbol": "Messi", "tenis": "Nadal", "basket": "Jordan"}

	// for key, value := range deportes {
	// 	fmt.Println("Llave: ", key, "Valor: ", value)
	// } 
	calificaciones := map[string]int{"Juan": 10, "Pedro": 9, "Maria": 8}
	for key, value := range calificaciones {
		fmt.Println("Nombre: ", key, "Calificación: ", value)
	}
}