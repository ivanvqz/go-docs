package main
import "fmt"

func main()  {
	var tablas int = 2
	for i := 1; i <= 10; i++ {
		// fmt.Printf("%d x %d = %d\n", tablas, i, tablas * i)
		fmt.Println(tablas, "x", i, "=", tablas * i)
	}
}