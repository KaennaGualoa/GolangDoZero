package main

import "fmt"

func main() {

	fmt.Println(soma(42, 13))

	sub := subtracao(10, 5)
	fmt.Println(sub)

	nome1, nome2, sobrenome := printaNome("Kaenna", "Gualoa")
	fmt.Println(nome1)
	fmt.Println(nome2)
	fmt.Println(sobrenome)
}

func subtracao(x int, y int) int {
	//ou
	//somaDosValores := x + Y
	//return somaDosValores
	return x - y
}

func soma(x int, y int) int {
	return x + y
}

//Função começando com a letra minúscula:
//Função é Privada
//Função ela só pode ser utilizada no próprio pacote
// Se for Maiuscula
// A função é publica e pode ser utilizada fora do proprio pacote
//Como utilizaria ela fora do pacote: main.PrintaNome(nome)
func printaNome(nome, sobrenome string) (string, string, string) {
	return nome, nome, sobrenome
}
