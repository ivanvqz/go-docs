package main
import "fmt"

func main()  {
	// Similar al ciclo while en otros lenguajes de programación
	var tablas int = 2
	i := 1
	for i <= 10{
		// fmt.Printf("%d x %d = %d\n", tablas, i, tablas * i)
		fmt.Println(tablas, "x", i, "=", tablas * i)
		i++
	}
}


// package main
// import "fmt"

// func main()  {}