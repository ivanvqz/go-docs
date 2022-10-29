package main

import "fmt"

func main() {
	type persona struct {
		Name string
		LastName string
		Age int
	}

	db := persona {
		Name: "Ivan",
		LastName: "Vazquez",
		Age: 23,
	}

	// fmt.Printf("%+v\n", db)

	// Acceder a los campos de la estructura
	fmt.Println("Nombre: ", db.Name)
	fmt.Println("Nombre: ", db.LastName)
	fmt.Println("Nombre: ", db.Age	)

	//Estructura literal
	git := persona {
		"Naomi",
		"Rosales",
		18,
	}
	otro := persona{Name: "Totoro"}
	fmt.Printf("%+v\n", git)

	fmt.Printf("%+v\n", otro)

	//punteros a estructuras
	p := &db

	//cambiar valores desde el puntero
	// (*p).LastName = "Velazquez"
	p.LastName = "Velazquez"
	fmt.Printf("%+v", db)
	fmt.Printf("%+v", p)

	//acceso a campos de una estructura

}