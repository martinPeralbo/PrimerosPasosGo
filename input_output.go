package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	edad := 22

	fmt.Printf("Hola mundo %d", edad)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ingresa un nombre")

	text, err := reader.ReadString('\n')

	if err != nil {

		fmt.Println(err)

	} else {
		fmt.Println("Hola" + text)
	}
}
