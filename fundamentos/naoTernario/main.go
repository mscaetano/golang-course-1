package main

import "fmt"

//nao tem operador ternario no go

func resultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println("Situacao:", resultado(6.2))
}
