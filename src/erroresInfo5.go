package main

import (
	"fmt"
	"log"
)

type ubicacionError struct {
	lat  string
	long string
	err  error
}

func (n ubicacionError) Error() string {
	return fmt.Sprintf("Un error de concepto matemático ha ocurrido en: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("\nNo hay raíz cuadrada real de un número negativo: %v.", f)
		return 0, ubicacionError{"50.2229 N", "99.4656 W", nme}
	}
	return 42, nil
}
