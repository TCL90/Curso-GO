package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println()

	wg.Add(2)

	go goroutine1()
	go goroutine2()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println()

	wg.Wait()

	fmt.Println()
	fmt.Println("Finaliza main.")
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}

func goroutine1() {
	defer wg.Done()
	fmt.Println("Esta es la Goroutine1.")
}

func goroutine2() {
	defer wg.Done()
	fmt.Println("Esta es la Goroutine2.")
}
