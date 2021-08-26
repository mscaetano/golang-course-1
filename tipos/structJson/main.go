package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct to json
	p1 := produto{1, "Notebook", 1899.90, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1) //retorna um slice de byte
	fmt.Println(string(p1Json))   //converto o slice de byte para uma string

	// json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta","preco":8.90,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2) //converto o primeiro parametro para um slice de byte
	fmt.Println(p2.Tags[1])
	fmt.Println(reflect.TypeOf(p2))
	fmt.Println(p2.ID)
}
