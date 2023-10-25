package average

import "fmt"

func main() {
	fmt.Println("La media de 1 y 2 es:", media([]int{1, 2}...))
	fmt.Println("La media de 1, 2, 3, 4, 5, 6, 7, 8 es:", media([]int{1, 2, 3, 4, 5, 6, 7, 8}...))
	fmt.Println("La media de 11, 12 es:", media([]int{11, 12}...))
	fmt.Println("La media de 34, 25 es:", media([]int{34, 25}...))
	fmt.Println("La media de 14, 56, 85, 32, 12 es:", media([]int{14, 56, 85, 32, 12}...))
}

func media(xi ...int) float64 {
	var resultado int
	for _, v := range xi {
		resultado += v
	}
	return float64(resultado) / float64(len(xi))
}
