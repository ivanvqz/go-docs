package main

import (
	"fmt"
	"os"
)

func main()  {
	// En la vida real funciona como en casos de limpieza de recursos, cerrar conexiones de red, cerrar archivos o controladores de datos
	file, err := os.Create("prueba.txt")
	if err != nil {
		fmt.Printf("Ocurrió un error al crear el archivo: %v", err)
		return //para que se detenga el programa
	}

	_, err = file.Write([]byte("Hola mundo desde un archivo."))
	if err != nil {
		fmt.Printf("Ocurrió un error al escribir el archivo: %v", err)
		return //para que se detenga el programa
	}
	file.Close()
}