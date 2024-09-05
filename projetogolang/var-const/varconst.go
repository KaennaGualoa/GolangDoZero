package main

import "fmt"

func main() {
	//variaveis
	//var + nome da variavel + tipo
	var nome string
	nome = "Maria"
	fmt.Println(nome)

	nome = "Kaenna"
	fmt.Println(nome)

	var idade int
	idade = 28
	fmt.Println(idade)

	var a = "Kaenna"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	f := "Apple"
	fmt.Println(f)

	k, m := "Kaenna", "Maria"
	fmt.Println(k, m)

	i := 524
	fmt.Println(i)

	//const nao muda
	const idadeMaria = 10
	fmt.Println(idadeMaria)

}
