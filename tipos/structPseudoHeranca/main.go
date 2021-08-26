package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       //composicao -> nao coloco o identificador apenas o tipo -> campo anonimo
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "f.40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s está com turbo ligado? %v\n", f.nome, f.turboLigado)
	fmt.Println(f)
}
