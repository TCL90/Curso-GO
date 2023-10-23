package main

import (
	"fmt"
)

func main() {
	const (
		_ = iota
		_
		_
		a = iota + 2020
		b = iota + 2020
		c = iota + 2020
	)

	fmt.Println(a, b, c)
}