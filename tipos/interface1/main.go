package main

import (
	"fmt"
	"reflect"
)

type imprimivel interface {
	toString() string //metodo relacionado a struct para ser compativel com a interface
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string { //metodo toString da interface - logo permite a struct ser do tipo imprimivel
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string { //metodo toString da interface
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Roberto", "Silva"} //pessoa tem o metodo implementado pela interface imprimivel por isso a atribuicao - logo permite a struct ser do tipo imprimivel
	fmt.Println(coisa.toString())

	imprimir(coisa)

	coisa = produto{"Calça Jeans", 79.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	p2 := produto{"Calça Jeans", 179.90}
	fmt.Println(reflect.TypeOf(p2))
	imprimir(p2) // metodo recebe uma var do tipo imprimivel - produto respeita a implementacao de todos os metodos da interface imprimivel, por isso podemos passar ele
}
