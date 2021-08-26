package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

// go test -cover
// go teste -coverprofile=resultado.out -> gerar aquivo detalhado
// go toll cover -func=resultado.out -> ver resultado do arquivo
// go tool cover -html=resultado.out
