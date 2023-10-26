package mystring

import (
	"fmt"
	"strings"
	"testing"
)

func TestConcat(t *testing.T) {
	s := "CACC, Catch As Catch Can"
	xs := strings.Split(s, " ")
	s = Concat(xs)
	if s != "CACC, Catch As Catch Can" {
		t.Error("Got", s, "wanted", "CACC, Catch As Catch Can")
	}
}

func TestJoin(t *testing.T) {
	s := "CACC, Catch As Catch Can"
	xs := strings.Split(s, " ")
	s = Join(xs)
	if s != "CACC, Catch As Catch Can" {
		t.Error("Got", s, "wanted", "CACC, Catch As Catch Can")
	}
}

func ExampleConcat() {
	s := "La reunión empieza en 45 minutos"
	xs := strings.Split(s, " ")
	fmt.Println(Concat(xs))
	// Output:
	// La reunión empieza en 45 minutos
}

func ExampleJoin() {
	s := "La reunión empieza en 45 minutos"
	xs := strings.Split(s, " ")
	fmt.Println(Join(xs))
	// Output:
	// La reunión empieza en 45 minutos
}

func BenchmarkConcat(b *testing.B) {
	s := "CACC, Catch As Catch Can"
	xs := strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = Concat(xs)
	}
}

func BenchmarkJoin(b *testing.B) {
	s := "CACC, Catch As Catch Can"
	xs := strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = Join(xs)
	}
}
