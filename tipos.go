package main

import (
	"fmt"
	"strconv"
)

func main() {

	edad := 22

	edad_str := strconv.Itoa(edad)
	edad_int, _ := strconv.Atoi(edad_str)
	fmt.Println(edad_int + 10)

}
