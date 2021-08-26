package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependent(t *testing.T) {
	t.Parallel() //executar em paralelo - ganho de tempo
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}

	t.Fail()
}

//dentro da pasta teste execitar:
//go test ./.. -> rodando todos os testes do pacote
