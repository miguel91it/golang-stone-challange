package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestObterJogadores(t *testing.T) {

	t.Run("retornar resultado de Maria", func(t *testing.T) {

		requisicao, _ := http.NewRequest(http.MethodGet, "/jogadores/Maria", nil)

		resposta := httptest.NewRecorder()

		ServidorJogador(resposta, requisicao)

		recebido := resposta.Body.String()
		esperado := "20"

		if recebido != esperado {
			t.Errorf("recebido '%s', esperado '%s'", recebido, esperado)
		}
	})
}
