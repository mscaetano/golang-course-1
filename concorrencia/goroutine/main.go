package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Pq vc não fala comigo?", 3)
	// fale("João", "Só posso falar depois de vc!", 1)
	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Opa...", 500)
	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns!", 5)

	// executa a funcao independente até a
	// go routine principal - main, do programa
	// estar executando, ou seja ela pode finalizar
	// antes mesmo da goroutine finalizar todo o processo, POR ENQUANTO
}
