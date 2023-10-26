package mystring2

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("James")
	if s != "Hello my dear, James" {
		t.Error("Got", s, "wanted", "Hello my dear, James")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Hello my dear, James
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
