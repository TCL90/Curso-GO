package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c
	fmt.Println(ok, v)

	v, ok = <-c
	fmt.Println(ok, v)

	fmt.Println("Finalizado.")
}
