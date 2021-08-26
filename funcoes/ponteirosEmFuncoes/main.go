package main

import "fmt"

func ini1(n int) {
	n++
}

func ini2(n *int) {
	*n++
}

func main() {
	n := 1

	ini1(n) //por valor - nao tem modificacao
	fmt.Println(n)

	//& -> passando o endereco de memoria da variavel
	ini2(&n) // por referencia - tem modificacao
	fmt.Println(n)
}
