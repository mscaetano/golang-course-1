package main

import "fmt"

func closure() func() {
	x := 10

	var funcao = func() {
		fmt.Println(x)
	}

	return funcao
}

func main() {
	x := 20

	fmt.Println(x)

	imprimeX := closure()
	imprimeX()

}

// package main

// import "fmt"

// func adder() func(int) int {
// 	sum := 0
// 	return func(x int) int {
// 		sum += x
// 		return sum
// 	}
// }

// func main() {
// 	pos, neg := adder(), adder()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(
// 			pos(i),
// 			neg(-2*i),
// 		)
// 	}
// }
