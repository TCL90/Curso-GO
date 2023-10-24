package main

import "fmt"

func main() {
	c := make(chan int)

	escribe(c)
	recibe(c)
}

func escribe(cs chan<- int) {
	go func() {
		for i := 0; i < 100; i++ {
			cs <- i
		}
		close(cs)
	}()
}

func recibe(cr <-chan int) {
	for v := range cr {
		fmt.Println(v)
	}
}
