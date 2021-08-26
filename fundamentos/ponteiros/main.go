package main

import "fmt"

func main() {
	i := 1

	//Go nao tem aritmetica de ponteioros
	var p *int = nil //iniciei com o nulo

	p = &i //pegar o endereco da variavel e atribuo para o ponteiro - deste modo guarda o endereco de memoria da variavel i
	*p++   // desreferenciado o ponteiro | peguei o valor associado ao ponteiro
	i++

	fmt.Println(p, *p, i, &i)

}
