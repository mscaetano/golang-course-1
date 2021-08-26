package main

import "fmt"

func main() {

	funcionarios := map[string]float64{
		"Maria":     1500.00,
		"Jozao":     1500.50,
		"Terezinha": 1700.60,
	}

	funcionarios["Raphael"] = 1350.00
	delete(funcionarios, "Jozao")

	fmt.Println("Nome      Salario")

	//primeira sempre é a chave do map e a segunda é o valor
	for nome, salario := range funcionarios {
		fmt.Println(nome, "     ", salario)
	}

}
