package main

import (
	"log"
)

var (
	// unico requisito deste slice é que os valores das notas já estejam
	// ordernados de forma decrescente. Uma forma elegante de garantir isto
	// sempre poderia ser ao subir o serviço, porém os testes irão cobrir
	// alterações erroneas. Alterar as possíveis notas vai afetar o resultados
	// dos testes também
	Notas = []uint64{50, 10, 5, 1}
)

func DevolverNotas(valorTotal uint64) map[uint64]uint64 {
	//
	notasParaDevolver := map[uint64]uint64{}

	// após cada iteração, vamos subtrair o valor restante
	valorRestante := valorTotal

	posSlice := 0
	for {
		//
		nota := Notas[posSlice]

		if nota <= valorRestante {
			valorRestante -= nota
			notasParaDevolver[nota]++
		} else {
			posSlice++
		}

		if valorRestante <= 0 {
			break
		}
	}

	log.Printf("DevolverNotas recebeu [%d] e devolveu [%+v]", valorTotal, notasParaDevolver)

	return notasParaDevolver
}
