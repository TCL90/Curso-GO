package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Nº CPUs", runtime.NumCPU())
	fmt.Println("Nº Goroutines", runtime.NumGoroutine())

	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)
	contador := 0
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
		}()
		fmt.Println("Nº Goroutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Cuenta", contador)
}
