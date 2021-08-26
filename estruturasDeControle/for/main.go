package main

import (
	"fmt"
	"time"
)

func main() {

	//funciona como while
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// outro tipo de for

	for i := 0; i <= 20; i += 2 {
		fmt.Print(i, " ")
	}

	// teste par ou impar

	for i := 1; i < 11; i++ {
		if i%2 == 0 {
			fmt.Println(i, "- Par")
		} else {
			fmt.Println(i, "- Impar")
		}
	}

	for {
		//laco for pra sempre
		fmt.Println("Pra sempre...")
		time.Sleep(time.Second)
		//de 5 em 5 segundos printa - time.Sleep(time.Second * 5)
	}
}
