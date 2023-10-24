package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	const gorutinas int = 10
	wg.Add(gorutinas)
	c := make(chan int)
	for i := 0; i < gorutinas; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
			wg.Done()
		}()
	}

	for j := 0; j < 100; j++ {
		fmt.Println(<-c)
	}
	wg.Wait()
}
