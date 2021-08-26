package main

import (
	"fmt"
	"time"
)

func tipoParam(i interface{}) string {

	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "funcao"
	default:
		return "tipo nao identificado"
	}
}

func main() {
	fmt.Println(tipoParam(2.3))
	fmt.Println(tipoParam(1))
	fmt.Println(tipoParam("Opa"))
	fmt.Println(tipoParam(func() {}))
	fmt.Println(tipoParam(time.Now()))
}
