package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hola")
	fmt.Fprintln(os.Stdout, "Hola")
	io.WriteString(os.Stdout, "Hola")
}