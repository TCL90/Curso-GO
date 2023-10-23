package main

import (
	"fmt"
	"sort"
)

func main() {
	sliceInt := []int{9, 54, 23, 1, 6, 8}
	sliceString := []string{"Zen", "Alberto", "Mario", "Begoña", "Geraldo"}

	fmt.Println(sliceInt)
	fmt.Println(sliceString)

	sort.Ints(sliceInt)
	fmt.Println(sliceInt)

	sort.Strings(sliceString)
	fmt.Println(sliceString)
}