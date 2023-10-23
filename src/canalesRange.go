package main

import "fmt"

func main() {
	c := make(chan int)

	// Send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	// Receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Finalizado.")
}

func enviar(c chan<- int) {
	c <- 42
}

func recibir(c <-chan int) {
	fmt.Println(<-c)
}
