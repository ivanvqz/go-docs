package main
import (
	// Permite leer el contenido de un archivo
	// "io/ioutil"
	"fmt"
)

func main()  {
	// // Leer el contenido de un archivo
	// content, err := ioutil.ReadFile("./archivo.txt")
	// if err != nil { //si el error es diferente de nil(de su valor 0)
	// 	fmt.Printf("Error al leer el archivo: %v", err)
	// 	return
	// }

	// fmt.Println(string(content))
	resultado, err := division(100, 3)	
	if err != nil { //si el error es diferente de nil(de su valor 0)
		err = fmt.Errorf("Ocurrió un error: %v", err)
			return
		}
	fmt.Println(resultado)

}

// func division(num1, num2 int) float32{
// 	if num2 == 0 {
// 		panic("No se puede dividir entre cero")
// 	}
// 	return float32(num1) / float32(num2)
// }

func division(num1, num2 int) (resultado float32, err error){
	if num2 == 0 {
		return 0, fmt.Errorf("No se puede dividir entre cero")
	}
	return float32(num1) / float32(num2), nil
}