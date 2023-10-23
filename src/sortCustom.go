package main

import (
	"fmt"
	"sort"
)

type persona struct {
	Nombre string
	Edad   int
}

type PorEdad []persona

func (a PorEdad) Len() int           { return len(a) }
func (a PorEdad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PorEdad) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

type PorNombre []persona

func (a PorNombre) Len() int           { return len(a) }
func (a PorNombre) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PorNombre) Less(i, j int) bool { return a[i].Nombre < a[j].Nombre }


func main() {
	p1 := persona{"Tomás", 25}
	p2 := persona{"Laura", 80}
	p3 := persona{"Marta", 19}
	p4 := persona{"Manuel", 35}

	personas := []persona{p1, p2, p3, p4}

	fmt.Println(personas)

	sort.Sort(PorEdad(personas))
	fmt.Println(personas)

	sort.Sort(PorNombre(personas))
	fmt.Println(personas)
}
