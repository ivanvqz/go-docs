package main
import "fmt"

// func main()  {
// 	nums := []int{1,2,3,4,5,6,7,8,9,10}
// 	result := filter(nums, func(num int) bool {
// 		if num <= 10 {
// 			return true
// 		}

// 		return false
// 	})

// 	fmt.Println(result)
// }

// func filter(nums []int, callback func(int) bool) []int {
// 	resultado := []int{}
// 	for _, v := range nums {
// 		if callback(v){
// 			resultado = append(resultado, v)
// 		}
// 	}
// 	return resultado
// }

//Funciones que retornan funciones
func main()  {
	saluda := hola("Ivan")("Vazquez") // se ejecuta la segunda funcion

	fmt.Println(saluda)
}

func hola(name string) func(string) string{
	return func (apellido string) string {
		return "Hola " + name + " " + apellido
	}
}