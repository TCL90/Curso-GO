package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("Introduce un número entero: ")
	fmt.Scan(&num)
	fmt.Printf("En base decimal es %d, en binario %b y en hexadecimal %x", 
	num, num, num)
}