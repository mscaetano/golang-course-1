package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println("Essa var é do tipo: ", reflect.TypeOf(3200))

	//                             byte   short   int     long
	// sem sinal (só positivos)... uint8  uint16  uint32  uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais(float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é: ", reflect.TypeOf(x))
	fmt.Println("O tipo do literal do 49.99: ", reflect.TypeOf(49.99)) //float64

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Olá meu nome é Leo"
	fmt.Println(s1 + "!")                         //concatenacao
	fmt.Println("O tamanho da string é", len(s1)) //tamanho da string - 20

	// string com multiplas linhas
	s2 := `Olá meu
	nome
	é
	Leo`
	fmt.Println("O tamanho da string é", len(s2)) //tamanho da string - 24 - devidos as quebras de linha

	// no go nao temos tipo char!!!

}
