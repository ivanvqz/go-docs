package main
import "fmt"

func main()  {
	// para cantidades indefinidas de datos
	// se usa el operador "..."
	suma(2, 4 ,23, 342, 234, 123)
}

func suma(numeros ...int)  {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	fmt.Println(total)
}