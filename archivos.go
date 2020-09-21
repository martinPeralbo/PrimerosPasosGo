package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	/*file_data, err := ioutil.ReadFile("./hola.txt")
	if err != nil {
		fmt.Print("Hubo un error")

	}

	fmt.Println(string(file_data))
	*/
	archivo, error := os.Open("./hola.txt")
	if error != nil {
		fmt.Println("Hubo un error")
	}
	defer print("Esto se escribe si o si")
	scanner := bufio.NewScanner(archivo)
	i := 1
	for scanner.Scan() {

		linea := scanner.Text()
		fmt.Print(i)
		fmt.Println(" " + linea)
		i++
	}

	archivo.Close()
}
