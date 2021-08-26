package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return //return limpo -> ja que eu ja populei as minhas variaveis de retorno -> retorna na ordem que coloquei entre parenteses
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)
}
