package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	//slice1 = append(array1[:], 10, 9) //pegou tudo do array1
	slice1 = append(slice1, 4, 5, 6) //adiciona elementos dentro de um slice
	slice1 = append(slice1, 7)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1) // pega os dois primeiros elementos do slice1 - 4,5 - no copy sempre tem que ser um slice
	fmt.Println(slice2)

	var slice []int
	fmt.Println(slice)
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)
	for _, num := range slice2 {
		slice = append(slice, num)
	}
	fmt.Println(slice)
}
