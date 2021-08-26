package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Numero invalido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, _ := fatorial(3)
	fmt.Println(resultado)

	_, err := fatorial(-4)

	if err != nil {
		fmt.Println(err)
	}
}

//solucao mais simples, permitindo o usuario passar apenas numero positivos(sem sinal) - uint
// package main
// import "fmt"
// func fatorial(n uint) uint { switch {
// case n == 0: return 1
// default:
// return n * fatorial(n-1) }
// }
// func main() {
// resultado := fatorial(5)
// fmt.Println(resultado) }
