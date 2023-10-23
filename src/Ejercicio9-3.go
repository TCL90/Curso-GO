package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var contador int = 0
	var wg sync.WaitGroup

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println()
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			incremento := contador
			runtime.Gosched()

			incremento++
			contador = incremento
			wg.Done()
			fmt.Println("Incrementador:", incremento)
		}()
	}

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println()
	wg.Wait()

	fmt.Println("Contador Final:", contador)
}
