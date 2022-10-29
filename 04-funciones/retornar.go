package main
import "fmt"
func main()  {
	resultado := suma(213 ,100)
	fmt.Println(resultado)
}
//cuando dos o más parámetros tienen un mismo tipo, se puede omitir el tipo de los parámetros posteriores
func suma(num, num2 int)int {
	return num + num2
}

//el ultimo int es el valor que se va a retornar