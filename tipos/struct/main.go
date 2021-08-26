package main

import "fmt"

type produto struct {
	nome      string
	preco     float64
	desconto  float64
	categoria string
}

// metodo  : funcao com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) getNome() string {
	return p.nome
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:      "Lapis",
		preco:     1.79,
		desconto:  0.05,
		categoria: "material escolar",
	}

	fmt.Println(produto1)
	fmt.Printf("%.2f", produto1.precoComDesconto())

	produto2 := produto{"Notebook", 3000.99, 0.10, "Eletronicos"}
	fmt.Println(produto2)

	var pessoa1 pessoa
	pessoa1 = pessoa{
		nome:  "TesteNome",
		idade: 20,
	}
	fmt.Println(pessoa1)

	var nomePessoa2 string
	var idadePessoa2 int

	nomePessoa2 = "testePessoa2"
	idadePessoa2 = 21

	pessoa2 := pessoa{nomePessoa2, idadePessoa2}
	fmt.Println(pessoa2)
	fmt.Println(pessoa2.getNome())
}
