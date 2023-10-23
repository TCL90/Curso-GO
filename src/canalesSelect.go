package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan int)

	// Enviar
	go enviar(par, impar, salir)

	// Recibir
	recibir(par, impar, salir)

	fmt.Println("Finalizado")
}

// Enviar
func enviar(p, i, s chan<- int) {
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			p <- j
		} else {
			i <- j
		}
	}

	close(p)
	close(i)

	s <- 0
}

func recibir(p, i, s <-chan int) {

}
