package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("Contexto:\t\t", ctx)
	fmt.Println("Error de Contexto:\t", ctx.Err())
	fmt.Printf("Tipo de Contexto:\t%T\n", ctx)

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println()
	fmt.Println("Contexto:\t\t", ctx)
	fmt.Println("Error de Contexto:\t", ctx.Err())
	fmt.Printf("Tipo de Contexto:\t%T\n", ctx)
	fmt.Println("Cancel Func:\t\t", cancel)
	fmt.Printf("Tipo de Cancel:\t\t%T\n", cancel)
	fmt.Println()

	cancel()

	fmt.Println("Contexto:\t\t", ctx)
	fmt.Println("Error de Contexto:\t", ctx.Err())
	fmt.Printf("Tipo de Contexto:\t%T\n", ctx)
	fmt.Println("Cancel Func:\t\t", cancel)
	fmt.Printf("Tipo de Cancel:\t\t%T\n", cancel)
}
