package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	result := soma(1, 2)
	fmt.Println(result)

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(5, 3))
}
