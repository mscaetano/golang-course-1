package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//map devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[1] = "Maria"
	aprovados[2] = "Jozao"
	aprovados[3] = "Carlos"
	aprovados[4] = "teste"

	fmt.Println("ID      Nome")
	for id, nome := range aprovados {
		fmt.Println(id, "     ", nome)
	}

	fmt.Println(aprovados[1]) //acessar determinado valor do map pela chave
	delete(aprovados, 4)      // deletar passo o map e a chave a ser excluida

	fmt.Println("ID      Nome")
	for id, nome := range aprovados {
		fmt.Println(id, "     ", nome)
	}
}
