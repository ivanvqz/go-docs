package main

import "fmt"

func main() {
	animales := make(map[string]string)
	animales["perro"] = "woof"
	animales["gato"] = "meow"
	animales["caballo"] = "neigh"
	fmt.Println("Animales en inglés: \n", animales)

	//mapa literal
	frutas := map[string]string{
		"manzana": "apple",
		"naranja": "orange",
		"platano": "banana",
		"kaki":    "persimmon",
	}
	fmt.Println("Frutas: \n", frutas)

	//eliminar elementos
	// Primer argumento: la clave del elemento que queremos eliminar
	// segundo argumento: la llave que queremos eliminar
	delete(frutas, "kaki")
	fmt.Println("Frutas sin el kaki: \n", frutas)

	//obtener elementos
	animales["perro"] = "woof"
	fmt.Println("¿Cómo hace el perro?: ", animales["perro"])

	//comprobar si existe una clave
	// Primer argumento: el mapa en el que queremos buscar
	// segundo argumento: la llave que queremos buscar
	_, existe := animales["gato"]
	fmt.Println("¿Existe el gato?: ", existe)

	//recorrer un mapa
	for llave, valor := range animales {
		fmt.Println("El", llave, "dice", valor)
	}

	//añadir elementos inexsistentes
	if animal, ok := animales["gatito"]; !ok {
		animales["gatito"] = "nya"
		fmt.Println("El gatito dice", animal)
	}
	fmt.Println("Animales: \n", animales)
}