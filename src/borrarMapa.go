package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Batman": 32,
		"Robin": 27
	}

	fmt.Println(m)

	delete(m, "Batman")
	fmt.Println(m)

	if v, ok := m["Tomás"]; ok {
		delete(m, "Tomás")
	}
	fmt.Println(m)

}