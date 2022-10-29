package main

import "fmt"

func main() {
	diasSemana := 7
	switch diasSemana {
	case 1:
		fmt.Println("Lunes")
	
	case 2:
		fmt.Println("Martes")
	
	case 3:
		fmt.Println("Miercoles")
	
	case 4:
		fmt.Println("Jueves")
	
	case 5:
		fmt.Println("Viernes")
	
	//se pueden agrupar algunos casos
	case 6, 7:
		fmt.Println("A descansar!!!")

	default:
		fmt.Println("Error")
	}

	edad := 56
	switch {
	case edad >=18 && edad <= 25:
		fmt.Println("Aceptado, entra")

	default:
		fmt.Println("No eres acepado")
	}
}