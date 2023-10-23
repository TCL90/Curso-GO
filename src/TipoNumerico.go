package main

import "fmt"
import "runtime"

var z uint16 // Si ponemos uint8 se desborda, con uint16 aguanta 256

func main() {
	x := 42
	y := 42.32
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	z = 256
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	z = uint16(x)
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(runtime.GOOS)
	fmt.PRintln(runtime.GOARCH)
}