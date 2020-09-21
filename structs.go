package main

import "fmt"

//Clase
type User struct {
	edad             int
	nombre, apellido string
}

func main() {

	var martin = User{nombre: "pepe", apellido: "fernan", edad: 30}
	var jose = new(User)

	fmt.Println(User{nombre: "Uriel", apellido: "Fernandez"})
	fmt.Println(martin)
	fmt.Println(jose)
	fmt.Println(*jose)
}
