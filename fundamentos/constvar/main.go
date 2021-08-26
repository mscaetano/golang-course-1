package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	// forma reduzida de criar uma var
	// coloquei o m no import e posso utilizar o nome do import dado
	area := PI * m.Pow(raio, 2)
	fmt.Println("Area:", area)

	// outra forma de declarar constantes e variaveis
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	g, h, i := 2, true, "Teste"

	fmt.Println(e, f, g, h, i)
}
