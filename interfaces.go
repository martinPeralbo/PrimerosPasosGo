package main

import "fmt"

type User interface {
	Permisos() int
	Name() string
}

type Admin struct {
	name string
}

func (this Admin) Permisos() int {
	return 5
}

func (this Admin) Name() string {
	return this.name
}
func auth(user User) string {

	if user.Permisos() > 4 {
		return user.Name() + " tiene permisos the administrador"

	} else if user.Permisos() == 4 {

		return user.Name() + " tiene permisos the editor "
	}

	return ""

}

type Editor struct {
	name string
}

func (this Editor) Permisos() int {
	return 4
}

func (this Editor) Name() string {
	return this.name
}

func main() {

	admin := Admin{"Pepe"}
	edit := Editor{"Luis"}
	fmt.Println(auth(admin))
	fmt.Println(auth(edit))

}
