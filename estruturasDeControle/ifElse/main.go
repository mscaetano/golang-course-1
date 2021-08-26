package main

import "fmt"

func imprimirresultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado \nnota:", nota)
	} else {
		fmt.Println("Reprovado \nnota:", nota)
	}
}

func main() {
	imprimirresultado(7.3)
	imprimirresultado(5.3)
}
