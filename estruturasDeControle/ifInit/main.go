package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	// funcao para gerar numero aleatorio
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	//gera numero ate 10
	return r.Intn(10)
}

func main() {
	//conseguiu criar uma variavel instanciando dentro do if, ficando visivel só para o bloco do if
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou")
	} else {
		fmt.Println("Ganhou")
	}
}
