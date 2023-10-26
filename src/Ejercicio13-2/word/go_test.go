package word

import (
	"fmt"
	"quote"
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	s := "Hola soy Tomás"
	if Count(s) != 3 {
		t.Error("Expected 3", "Got", s)
	}
}

func TestUseCount(t *testing.T) {
	s := "Hola soy Tomás"
	m := make(map[string]int)
	m["Hola"] = 1
	m["soy"] = 1
	m["Tomás"] = 1
	if !reflect.DeepEqual(UseCount(s), m) {
		t.Error("Expected", m, "Got", UseCount(s))
	}
}

func TestUseCount2(t *testing.T) {
	s := "Hola soy Tomás"
	m := UseCount(s)
	for k, v := range m {
		switch k {
		case "Hola":
			if v != 1 {
				t.Error("Expected 1, Got", v)
			}
		case "soy":
			if v != 1 {
				t.Error("Expected 1, Got", v)
			}
		case "Tomás":
			if v != 1 {
				t.Error("Expected 1, Got", v)
			}
		}
	}
}

func ExampleCount() {
	s := "Hola soy Tomás"
	fmt.Println(Count(s))
	// Output:
	// 3
}

func ExampleUseCount() {
	s := "Hola soy Tomás"
	fmt.Println(UseCount(s))
	// Output:
	// map[Hola:1 Tomás:1 soy:1]
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
