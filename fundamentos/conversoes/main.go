package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y)) //como o y nao é do tipo float64, e sim do tipo int, temos que converte-lo

	nota := 5.9
	notaFinal := int(nota)
	fmt.Println("Nota Final:", notaFinal)

	// esta conversao é do numero para a tabela ask
	//fmt.Println("Teste " + string(97))

	//int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	//string para int - retorna dois valores - um numero e um erro
	// _ - dando este nome para var ele ignora o erro e nao preciso utilizala
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	//string para bolean
	a, _ := strconv.ParseBool("true")
	if a {
		fmt.Println("Verdadeiro")
	}
}
