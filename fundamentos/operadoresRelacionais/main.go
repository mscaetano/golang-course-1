package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("As datas são iguais:", d1 == d2)
	fmt.Println("As datas são iguais:", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Matheus"}
	p2 := Pessoa{"Matheus"}
	fmt.Println("As pessoas são iguais:", p1 == p2) // ele ve o nome e nao a posicao da memoria, mas quando lidarmos com ponteiros ai sim ele ira ver a posicao da memoria

}
