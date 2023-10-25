package main

import "fmt"

type errorPer struct {
	err     error
	linea   int
	funcion string
}

func (e errorPer) Error() string {
	return fmt.Sprintf("El error %v se da en la línea %d, en la función %s", e.err, e.linea, e.funcion)
}

func main() {
	err := errorPer{fmt.Errorf("Error 29: Falta de argumentos de entrada"), 29, "foo()"}
	foo(err)
}

func foo(err error) {
	fmt.Println(err)

	// Assertion
	fmt.Println(err.(errorPer).funcion)
}
