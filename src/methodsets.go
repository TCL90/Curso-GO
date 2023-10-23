package main

import (
	"fmt"
)

type contacto struct {
	nombre   string
	telefono int
}

func (c contacto) DevuelveMovil() int {
	return c.telefono
}

func (c *contacto) DevuelveMovil2() int {
	return c.telefono
}

type datosPersonales interface {
	DevuelveMovil() int
}

type datosPersonalesPuntero interface {
	DevuelveMovil() int
	DevuelveMovil2() int
}

func main() {
	c1 := contacto{
		nombre:   "Tomás",
		telefono: 658310185,
	}

	infoMovil(&c1)
	infoMovil(c1)

	infoMovil2(&c1)
	// infoMovil3(c1), este no funciona porque al tener un método que solo recibe punteros, solo se pueden usar punteros en el methodset
}

func infoMovil(d datosPersonales) {
	fmt.Println(d.DevuelveMovil())
}

func infoMovil2(d datosPersonalesPuntero) {
	fmt.Println(d.DevuelveMovil2())
}

func infoMovil3(d datosPersonalesPuntero) {
	fmt.Println(d.DevuelveMovil())
}
