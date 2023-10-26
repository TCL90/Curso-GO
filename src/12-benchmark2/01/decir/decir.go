package decir

import "fmt"

// Saludar nos permite dar la bienvenida a una persona, cuyo nombre pasamos en un string.
func Saludar(s string) string {
	return fmt.Sprint("Bienvenido ", s)
}
