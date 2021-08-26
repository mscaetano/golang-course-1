package main

import "fmt"

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}

	//posso usar a funcao catetos pois ainda estou dentro do pacote
	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))
}
