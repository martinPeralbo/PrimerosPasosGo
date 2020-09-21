package main

func main() {

	//slice := make([]int, 3, 5)
	//print(cap(slice))

	var x *int

	entero := 5

	x = &entero

	*x = 9
	println(*x)
	print(x)

}
