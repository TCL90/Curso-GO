package main

import (
	"fmt"
	"sort"
)

type usuario struct {
	Nombre   string
	Apellido string
	Edad     int
	Dichos   []string
}

type usuarioPorEdad []usuario

func (u usuarioPorEdad) Len() int           { return len(u) }
func (u usuarioPorEdad) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }
func (u usuarioPorEdad) Less(i, j int) bool { return u[i].Edad < u[j].Edad }

type usuarioPorNombre []usuario

func (u usuarioPorNombre) Len() int           { return len(u) }
func (u usuarioPorNombre) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }
func (u usuarioPorNombre) Less(i, j int) bool { return u[i].Nombre < u[j].Nombre }

type usuarioPorApellido []usuario

func (u usuarioPorApellido) Len() int           { return len(u) }
func (u usuarioPorApellido) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }
func (u usuarioPorApellido) Less(i, j int) bool { return u[i].Apellido < u[j].Apellido }

func main() {
	u1 := usuario{
		Nombre:   "Eduar",
		Apellido: "Tua",
		Edad:     32,
		Dichos: []string{
			"Cachicamo diciéndole a morrocoy conchudo",
			"La mona, aunque se vista de seda, mona se queda",
			"Poco a poco se anda lejos",
		},
	}

	u2 := usuario{
		Nombre:   "Carlos",
		Apellido: "Pérez",
		Edad:     27,
		Dichos: []string{
			"Camarón que se duerme se lo lleva la corriente",
			"A ponerse las alpargatas que lo que viene es joropo",
			"No gastes pólvora en zamuro",
		},
	}

	u3 := usuario{
		Nombre:   "Che",
		Apellido: "López",
		Edad:     54,
		Dichos: []string{
			"Ni lava ni presta la batea",
			"Hijo de gato, caza ratón",
			"Más vale pájaro en mano que cien volando",
		},
	}

	usuarios := []usuario{u1, u2, u3}

	fmt.Println("Original:")
	for _, v := range usuarios {
		fmt.Println(v.Nombre, v.Apellido, v.Edad)
		for _, dicho := range v.Dichos {
			fmt.Printf("\t%s\n", dicho)
		}
	}
	fmt.Println()

	sort.Sort(usuarioPorNombre(usuarios))
	fmt.Println("Ordenados por Nombre:")
	for _, v := range usuarios {
		fmt.Println(v.Nombre, v.Apellido, v.Edad)
		sort.Strings(v.Dichos)
		for _, dicho := range v.Dichos {
			fmt.Printf("\t%s\n", dicho)
		}
	}
	fmt.Println()

	sort.Sort(usuarioPorApellido(usuarios))
	fmt.Println("Ordenados por Apellido:")
	for _, v := range usuarios {
		fmt.Println(v.Nombre, v.Apellido, v.Edad)
		sort.Strings(v.Dichos)
		for _, dicho := range v.Dichos {
			fmt.Printf("\t%s\n", dicho)
		}
	}
	fmt.Println()

	sort.Sort(usuarioPorEdad(usuarios))
	fmt.Println("Ordenados por Edad:")
	for _, v := range usuarios {
		fmt.Println(v.Nombre, v.Apellido, v.Edad)
		sort.Strings(v.Dichos)
		for _, dicho := range v.Dichos {
			fmt.Printf("\t%s\n", dicho)
		}
	}
	fmt.Println()
}
