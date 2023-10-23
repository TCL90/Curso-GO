package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var contador int64 = 0
	var wg sync.WaitGroup

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println()
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			fmt.Println("Incrementador:", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println()
	wg.Wait()

	fmt.Println("Contador Final:", contador)
}
