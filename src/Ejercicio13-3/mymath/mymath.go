// Package mymath nos permite realizar operaciones matemáticas como la media
package mymath

import "sort"

// CenteredAvg calcula la media de una lista de números después de eliminar
// los valores más grandes y los más pequeños.
func CenteredAvg(xi []int) float64 {
	sort.Ints(xi)
	xi = xi[1:(len(xi) - 1)]

	n := 0
	for _, v := range xi {
		n += v
	}

	f := float64(n) / float64(len(xi))
	return f
}
