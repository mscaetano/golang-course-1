package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 //ou exclusivo
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, false)
	fmt.Println("Tv 50: ", tv50)
	fmt.Println("Tv 32: ", tv32)
	fmt.Print("Sorvete: ", sorvete)

}
