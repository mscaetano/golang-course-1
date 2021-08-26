package main

import "fmt"

func main() {
	// compiladorr que ira contar e iniciara o array com o tamanho de valores que foi colocado nele
	// [...] contar o array a partir da quantidade elementos que foi inicializado o array
	numeros := [...]int{10, 20, 30, 40, 50}

	//range numeros - me manda dois atributos o i (indice) e o numero (valor do indice do vetor)
	for i, numero := range numeros {
		fmt.Printf("[%d] = %d\n", i, numero)
	}

	for _, numero := range numeros {
		fmt.Printf("%d\n", numero)
	}
}
