package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Subindo API na porta 8080")

	hNotas := &handlerNotas{}
	http.Handle("/notas", hNotas)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
