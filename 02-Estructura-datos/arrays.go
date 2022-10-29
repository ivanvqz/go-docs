package main

import "fmt"

func main() {

	// Create an array of 5 integers that is initialized to its zero value.
	var a [5]int
	fmt.Println("emp:", a)

	//en la posición 4 del array a, asigna el valor 100
	a[0] = 100
	a[1] = 200
	a[2] = 300
	a[3] = 400
	a[4] = 500
	fmt.Println("set:", a)
	fmt.Println("get:", a[3])//imprime el valor de la posición 4 del array a

	fmt.Println("len:", len(a))//imprime la longitud del array a

	//arrays literales
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//recorrer un array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			//asigna el valor de i+j a la posición i,j del array twoD
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}