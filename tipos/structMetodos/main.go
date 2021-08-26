package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	nomePartes := strings.Split(nomeCompleto, " ")
	p.nome = nomePartes[0]
	p.sobrenome = nomePartes[1]
}

func main() {
	pessoa1 := pessoa{"Pedro", "Silva"}
	fmt.Println(pessoa1.getNomeCompleto())

	pessoa1.setNomeCompleto("Maria Pereira")
	fmt.Println(pessoa1)
}
