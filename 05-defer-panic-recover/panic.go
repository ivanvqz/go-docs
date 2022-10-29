package main
import "fmt"

func main()  {
	//permite finalizar el programa cuando la función es llamada
	division(6, 2)
	division(46, 6)
	division(5, 0)
	division(4, 2)
	division(13, 13)
}

func division(a, b int) {
	defer func ()  {
		if r := recover(); r != nil {
			// únicamente se va a ejecutar si se produce un panic
			fmt.Println("Recuperándome del pánico: ", r)
		}
	}()
	validarDivision(b)
	fmt.Println(a / b)
}

func validarDivision(b int) {
	if b == 0 {
		panic("No se puede dividir por cero") 
	}
}

//recover: es posible recuperarse de la generación de un pánico