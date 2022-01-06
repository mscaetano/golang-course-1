package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	s := a[1:4]
	fmt.Println(s)

	var valor string = "valor"

	var ponteiro *string = nil

	fmt.Println(valor)
	fmt.Println(&valor)

	ponteiro = &valor

	fmt.Println(ponteiro)
	*ponteiro = "abc"
	fmt.Println(valor)
}
