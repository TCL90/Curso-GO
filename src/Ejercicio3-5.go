package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("Si dividimos %d entre 4, obtenemos el resto %d\n", i, i % 4)
	}
}