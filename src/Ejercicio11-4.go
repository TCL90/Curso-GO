package main

import (
	"fmt"
	"log"
)

type raizError struct {
	lat  string
	long string
	err  error
}

func (re raizError) Error() string {
	return fmt.Sprintf("error matemático: %v %v %v", re.lat, re.long, re.err)
}

func main() {
	_, err := raiz(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func raiz(f float64) (float64, error) {
	if f < 0 {
		return 0, raizError{"50.2289 N", "99.4656 W", fmt.Errorf("El Error está presente al usar el valor %v.", f)}
	}
	return 42, nil
}
