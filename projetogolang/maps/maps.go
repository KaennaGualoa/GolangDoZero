package main

import (
	"fmt"
)

func main() {
	idade := map[string]int{}
	idade["Kaenna"] = 27
	idade["bento"] = 4
	// fmt.Println(idade)
	// fmt.Println(idade["Kaenna"])
	// fmt.Println(idade["bento"])

	anoNasc := map[string]int{
		"Kaenna": 1996,
		"bento":  2008,
	}
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["Kaenna"])
	fmt.Println(anoNasc["bento"])
	anoNasc["golangDoZero"] = 2024
	fmt.Println(anoNasc)

	formacaoNota := map[string]float32{
		"Ciencia de Dados":          7.01,
		"Analise e Desenvolvimento": 8.01,
	}
	fmt.Println(formacaoNota)
	fmt.Println(formacaoNota["Analise e Desenvolvimento"])

	notaAprovacao := map[float64]string{
		5.01: "Reprovado",
		8.01: "Aprovado",
	}
	fmt.Println(notaAprovacao)
}

// 2 - Maps: Heterogêneos
// pode misturar tipos
// estrutura chave - valor
// [key] = value
// chave tem um tipo e o valor pode ter outro
// map[k]v -> k = chave, v = valor

// map[string]int
// {"steph": 27, "bento": 4}
// map[string]string
// {"steph": "cardoso", "bento": "cardoso"}
