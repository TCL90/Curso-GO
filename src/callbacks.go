package main

import "fmt"

func main() {
	sliceEnteros := []int{1, 2, 3, 4, 5, 6, 7, 8}

	total := suma(sliceEnteros...)
	fmt.Println(total)

	totalPares := sumaPares(suma, sliceEnteros...)
	fmt.Println(totalPares)

	totalImpares := sumaImpares(suma, sliceEnteros...)
	fmt.Println(totalImpares)
}

func suma(xi ...int) int {
	suma := 0
	for _, v := range xi {
		suma += v
	}

	return suma
}

func sumaPares(f func(xi ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}

func sumaImpares(f func(xi ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi {
		if v%2 != 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}
