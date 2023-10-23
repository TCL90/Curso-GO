package main

import "fmt"

func main() {
	var retornoInt int
	var retornoError error
	retornoInt, retornoError = fmt.Println("Hello World!", 42, true, 2.33)
	fmt.Println(retornoInt, retornoError)

	nb, er := fmt.Println("Hello World!", 42, true, 2.33)
	_, _ = fmt.Println(nb, er)
}
