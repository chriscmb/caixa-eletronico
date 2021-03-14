package main

import (
	"reflect"
	"testing"
)

func TestBasico(t *testing.T) {
	var testes = []struct {
		nomeTeste     string
		entrada       uint64
		saidaEsperada map[uint64]uint64
	}{
		{"Numero de notas para 93", 93, map[uint64]uint64{50: 1, 10: 4, 1: 3}},
		{"Numero de notas para 2", 2, map[uint64]uint64{1: 2}},
	}

	for _, st := range testes {
		// roda subtestes
		t.Run(st.nomeTeste, func(t *testing.T) {
			saidaRecebida := DevolverNotas(st.entrada)

			if !reflect.DeepEqual(st.saidaEsperada, saidaRecebida) {
				t.Errorf("O teste [%s] não retornou as notas esperadas. Retornou [%+v] quando deveria ter retornado [%+v]", st.nomeTeste, saidaRecebida, st.saidaEsperada)
			}
		})
	}
}
