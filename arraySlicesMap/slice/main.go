package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1))
	fmt.Println(reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	//slice nao é um array! Slice define um pedaco de um array
	s2 := a2[1:3] //vai do indice 1 ao 3, sem incluir o elemento de indice 3
	fmt.Println(s2)

	s3 := a2[:2] //vai do inicio do arry a2 ao indice 2 sem inclui-lo
	fmt.Println(s3)

	//posso fazer um slice de um slice
	s4 := s2[:1]
	fmt.Println(s2, s4)

	vetor := [...]int{1, 2, 3, 4, 5}
	sliceTeste := vetor[:4]
	fmt.Println(sliceTeste)

}
