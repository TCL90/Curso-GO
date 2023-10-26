package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	got := Years(7)
	if got != 49 {
		t.Error("Expected 49, Got", got)
	}
}

func TestYearsTwo(t *testing.T) {
	got := YearsTwo(7)
	if got != 49 {
		t.Error("Expected 49, Got", got)
	}
}

func ExampleYears() {
	fmt.Println(Years(7))
	// Output:
	// 49
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(7))
	// Output:
	// 49
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(7)
	}
}
