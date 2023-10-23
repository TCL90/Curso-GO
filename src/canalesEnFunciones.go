package main

import "fmt"

func main() {
	c := make(chan int)

	// Send
	go enviar(c)

	// Receive
	recibir(c)

	fmt.Println("Finalizado.")
}

func enviar(c chan<- int) {
	c <- 42
}

func recibir(c <-chan int) {
	fmt.Println(<-c)
}
