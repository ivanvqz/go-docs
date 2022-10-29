package main
import "fmt"

func main()  {
	//recorrer un string
	saludo := "Hola Mundo"
	for _, value := range saludo {
		fmt.Println(string(value))

		// *NOTA: el value muestra solamente los valores vinarios de cada caracter
	}
}