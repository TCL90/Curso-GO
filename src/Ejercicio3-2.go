package main

import (
	"fmt"
)

func main() {
	for x := 65; x <= 90; x++ {
		fmt.Printf("%d\n\t%#U\n\t%#U\n\t%#U\n\n", x, x, x, x)
	}

	for x := 65; x <= 90; x++ {
		fmt.Println(x)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", x)
		}
	}
}