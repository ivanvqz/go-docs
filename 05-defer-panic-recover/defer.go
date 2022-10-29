package main
import "fmt"

func main()  {
	// nos permite diferir(aplazar) algo hasta que se cumpla una condición
	// son como las pilas, se ejecutan en orden LIFO: Last In First Out, osea que el último 
	//que entra es el primero que sale
	// defer fmt.Println(3)
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// 2 1 3

	a := 5
	defer fmt.Println("Defer", a)
	a = 10
	fmt.Println("Normal", a)
}