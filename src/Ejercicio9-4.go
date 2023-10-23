package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var contador int = 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println()
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			incremento := contador

			incremento++
			contador = incremento
			mu.Unlock()
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
