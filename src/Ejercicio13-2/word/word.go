// Package word contiene utilidades para el conteo de palabras en un string base
package word

import "strings"

// UseCount cuenta las ocurrencias de cada palabra incluída en un string de entrada y
// las almacena en un map donde asocia cada string (key) con el número de ocurrencias
// (value) que se da en el string original pasado como argumento
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Cuenta el número de palabras total presentes en un string de entrada y lo devuelve
// como un int
func Count(s string) int {
	return len(strings.Fields(s))
}
