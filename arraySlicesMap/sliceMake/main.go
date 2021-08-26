package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) // capacidade de 20 elementos
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // ja tinha atingido a capacidade e eu coloquei mais um -> ele aumenta a capacidade (dobra)
	fmt.Println(s, len(s), cap(s))
}
