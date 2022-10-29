package main
import "fmt"

func main()  {
	// Este es un ciclo infinito, para detenerlo presiona Ctrl + C
	// puede ser ejecutados en procesos que se ejecuten en segundo plano, como sockets o hilos
	var tablas int = 2
	i := 1
	for{
		// fmt.Printf("%d x %d = %d\n", tablas, i, tablas * i)
		fmt.Println(tablas, "x", i, "=", tablas * i)
		i++
	}
}