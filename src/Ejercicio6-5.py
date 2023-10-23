package main

import (
	"fmt"
	"math"
)

type cuadrado struct {
	lado float64
}

type circulo struct {
	radio float64
}

func (c circulo) area() float64 {
	return c.radio * c.radio * math.Pi
}

func (c cuadrado) area() float64 {
	return c.lado * c.lado
}

type forma interface {
	area() float64
}

func info(f forma) {
	switch f.(type) {
	case circulo:
		fmt.Println("El área del círculo es", f.area())
	case cuadrado:
		fmt.Println("El área del cuadrado es", f.area())
	}
}

func main() {
	c1 := cuadrado{
		lado: 2.33,
	}

	c2 := circulo{
		radio: 3.2,
	}

	info(c1)
	info(c2)
}
