package main

import (
	"fmt"
	"mystring"
	"strings"
)

const s = "Nos preguntamos: ¿En que año fue 1 + 1?"

func main() {
	xs := strings.Split(s, " ")

	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", mystring.Concat(xs))
	fmt.Printf("\n%s\n", mystring.Join(xs))
}
