package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) // Receive-Only
	cs := make(chan<- int) // Send-Only

	fmt.Printf("Bidireccional:\t\t%T\n", c)
	fmt.Printf("Send-Only:\t\t%T\n", cs)
	fmt.Printf("Receive-Only:\t\t%T\n", cr)

	/*
		No se puede asignar los valores del especifico al general,
		es decir, no puedo asignar el valor de un Send-Only o un
		Receive-Only a un bidireccional.

		No funcionan:
		c = cr
		c = cs
	*/

	cr = c
	cs = c

	fmt.Println()
	fmt.Printf("Bidireccional a Send-Only\t\t%T\n", (<-chan int)(c))
	fmt.Printf("Bidireccional a Receive-Only\t\t%T\n", (chan<- int)(c))
}
