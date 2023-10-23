package main

import (
	"fmt"
	"x/crypto/bcrypt"
)

func main() {
	pass := `clave123`
	bs, err := bcrypt.GenerateFromPassword([]byte(pass), 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	claveLogin := `clave1234`
	err := bcrypt.CompareHashAndPassword(bs, []byte(claveLogin))
	if err != nil {
		fmt.Println("La contraseña que introduciste no es correcta.")
		return
	}
	fmt.Println("Accediste correctamente.")
}