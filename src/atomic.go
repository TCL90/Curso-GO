package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Nº CPUs", runtime.NumCPU())
	fmt.Println("Nº Goroutines", runtime.NumGoroutine())

	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)
	var contador int64 = 0

	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("Contador:", atomic.LoadInt64(&contador))

		}()
		fmt.Println("Nº Goroutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Cuenta", contador)
}
