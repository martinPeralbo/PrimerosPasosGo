package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	go mi_nombre_lento("Martin")
	fmt.Println("Que aburrido")
	fmt.Scanln()
}

func mi_nombre_lento(name string) {

	letras := strings.Split(name, "")

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}

}
