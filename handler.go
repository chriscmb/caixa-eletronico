package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type handlerNotas struct {
	Resposta stResp
}

type stResp struct {
	Mensagem string `json:"message"`
}

// lida exclusivamente com o método de notas
func (handler *handlerNotas) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("Requisiçao recebida")
	w.Header().Set("Content-Type", "application/json")

	// vamos permitir apenas o method GET para este endpoint
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(handler.Resposta.RetornoSimples("method inválido. Para este fonte de teste, apenas GET é suportado"))
		return
	}

	//
	valorNotasString := r.FormValue("valor")
	log.Printf("'valor' recebido: [%s]", valorNotasString)

	if valorNotasString == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(handler.Resposta.RetornoSimples("o campo 'valor' precisa ser informado"))
		return
	}

	//
	valorNotas, err := strconv.ParseUint(valorNotasString, 10, 64)
	if err != nil || valorNotas == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(handler.Resposta.RetornoSimples("o campo 'valor' aceita apenas numeros maiores que zero"))
		return
	}

	mapaNotasDevolver := DevolverNotas(valorNotas)

	w.WriteHeader(http.StatusOK)
	w.Write(handler.Resposta.RetornoSimples(fmt.Sprintf("%+v", mapaNotasDevolver)))
	return
}

func (st *stResp) RetornoSimples(s string) []byte {
	st.Mensagem = s

	retorno, err := json.Marshal(st)
	if err != nil {
		log.Fatalf("Não foi possível escrever um retorno para a requisião de API. Erro[%s]. Derrubando o serviço", err.Error())
	}
	return retorno
}
