// Package main es un ejercicio del bloque de benchmarking.
package mate

// Sum suma un conjunto ilimitado de ints y devuelve el resultado como un int
func Sum(xi ...int) int {
	var s int
	for _, v := range xi {
		s += v
	}

	return s
}
