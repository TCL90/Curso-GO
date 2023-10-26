package mystring

import "strings"

func Concat(xs []string) string {
	s := xs[0]
	for _, v := range xs[1:] {
		s += " "
		s += v
	}
	return s
}

func Join(xs []string) string {
	return strings.Join(xs, " ")
}
