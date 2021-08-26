package main

import "math"

// Inicando variavel com letra maiuscula é publico
// visibilidade dentro e fora do pacote
// nao posso ter funcoes duplicadas dentro do mesmo pacote

// Iniciar com letra minuscula é privado
// visivibilidade dentro do pacote

// Ex:
// Ponto - gerará algo publico
// ponto ou _Ponto - gerará algo privado

//ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é reponsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
