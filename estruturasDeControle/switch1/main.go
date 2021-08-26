package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)

	switch nota {
	case 10:
		//vai descer para o 9 e retornar do mesmo jeito, ou seja se for 10 e 9 e retorna A
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota invalida"
	}

}

func main() {
	fmt.Println("Conceito:", notaParaConceito(10))
	fmt.Println("Conceito:", notaParaConceito(7))
	fmt.Println("Conceito:", notaParaConceito(-1))
}
