package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonB = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll", "Order": "Dasyuromorphia"}
	]`)

	type Animal struct {
		Name  string
		Order string
	}

	var animals []Animal

	err := json.Unmarshal(jsonB, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}
