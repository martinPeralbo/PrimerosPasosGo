package main

type User struct {
	nombre, apellido string
}

type Tutor struct {
	Human
}

func main() {

	pepe := User{nombre: "Pepe", apellido: "fernan"}
	print(pepe.nombreCompleto())

	pTutor := Tutor{User{nombre: "juan"}}
	fmt.Println(pTutor.User.)
}

func (usuario User) nombreCompleto() string {

	nombre := usuario.nombre + " " + usuario.apellido
	return nombre
}
