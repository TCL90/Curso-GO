package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Check Error 1:", ctx.Err())
	fmt.Println("Num Gorutinas 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Trabajando", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Check Error 2:", ctx.Err())
	fmt.Println("Num Gorutinas 2:", runtime.NumGoroutine())

	fmt.Println("Cancelando Context.")
	cancel()

	time.Sleep(time.Second * 2)
	fmt.Println("Check Error 3:", ctx.Err())
	fmt.Println("Num Gorutinas 3:", runtime.NumGoroutine())
}
