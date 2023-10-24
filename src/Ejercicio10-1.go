package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	c2 := make(chan int, 2)

	go func() {
		c <- 42
	}()

	c2 <- 42
	c2 <- 43

	fmt.Println(<-c)
	fmt.Println(<-c2)
	fmt.Println(<-c2)
}
