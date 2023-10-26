package decir

import (
	"fmt"
	"testing"
)

func TestSaludar(t *testing.T) {
	s := Saludar("Tomás")
	if s != "Bienvenido Tomás" {
		t.Error("Esperaba", "Bienvenido Tomás", "obtuve", s)
	}
}

func ExampleSaludar() {
	fmt.Println(Saludar("Tomás"))
	//Output:
	//Bienvenido Tomás
}

func BenchmarkSaludar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Saludar("Tomás")
	}
}
