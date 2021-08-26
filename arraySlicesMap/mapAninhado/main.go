package main

import "fmt"

func main() {
	funcionarios := map[string]map[string]float64{
		"G": {
			"Gabriel Silva":       1500.00,
			"Gabriela dos Santos": 3000.50,
		},
		"H": {
			"Heitor Francisco": 4500.60,
			"Helena Silva":     2500.00,
		},
		"A": {
			"Ana Leticia": 4500.60,
		},
	}

	delete(funcionarios, "A")

	for letra, colaboradores := range funcionarios {
		fmt.Println(letra, colaboradores)
	}

}
