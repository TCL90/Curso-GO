package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan bool)

	go enviar(par, impar, salir)

	recibir(par, impar, salir)

	fmt.Println("Finalizado.")
}

// Canal para enviar
func enviar(par, impar chan<- int, salir chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(salir)
}

func recibir(par, impar <-chan int, salir <-chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("Desde el canal par", v)
		case v := <-impar:
			fmt.Println("Desde el canal impar", v)
		case i, ok := <-salir:
			if !ok {
				fmt.Println("Desde comma ok", i)
				return
			} else {
				fmt.Println("Desde comma ok", i)
			}
		}
	}
}
