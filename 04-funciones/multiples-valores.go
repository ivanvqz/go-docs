package main
import (
	"fmt"
	"strings"
)

func main()  {
	texto := "Me llamo Ivan"
	minu, mayu := convert(texto)
	fmt.Println(minu, mayu)
}

func convert(text string) (string, string) {
	min := strings.ToLower(text)
	may := strings.ToUpper(text)
	return min, may
}

//los parentesis son los calores a retirnar