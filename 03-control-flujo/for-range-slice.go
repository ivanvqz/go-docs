package main
import "fmt"

func main()  {
	nums := []uint8{2, 3, 4}
	// for i, v := range nums {
	// 	// se va a ejecutar 3 veces
	// 	fmt.Println("Indice",i, "Valor", v)
	// }
	for i := range nums {
		nums[i] *= 2
	}
	fmt.Println(nums)	
}