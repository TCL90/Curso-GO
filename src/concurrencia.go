package main

import (
	"fmt"
)

func haceAlgo(x int) int {
	return x * 5
}

func main() {
	ch := make(chan int)
	go func() {
		ch <- haceAlgo(5)
	}()
	fmt.Println(<-ch)
}