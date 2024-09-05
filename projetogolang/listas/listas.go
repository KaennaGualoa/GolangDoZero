package main

import "fmt"

func main() {

	//array - tamanho fixo
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	// fmt.Println(array[0])
	// fmt.Println(array[1])
	// fmt.Println(array[0], array[1])
	// fmt.Println(array)

	// numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	// fmt.Println(numPrimos)
	// fmt.Println(numPrimos[0:3]) // imprimir do 0 até o 2
	// fmt.Println(numPrimos[0])
	// fmt.Println(numPrimos[5])

	//Slices
	//var slice []string
	slice := make([]string, 5)
	slice[0] = "Hello"
	slice[1] = "People"
	// fmt.Println(slice[0], slice[1])
	// fmt.Println(slice[0])
	// fmt.Println(slice[1])
	// fmt.Println(slice[2])
	// slice[2] = "Kaenna"
	// fmt.Println(slice[2])
	// fmt.Println(slice)

	fmt.Println(len(slice))
	fmt.Println(slice[2])
	fmt.Println(slice[3])
	fmt.Println(slice[4])
	//fmt.Println(slice[5])

	numPares := []int{2, 4, 6, 8, 10, 12}
	fmt.Println(numPares)

	numPares = append(numPares, 14, 16, 19)
	fmt.Println(numPares)

}

//Listas

// 1 - Arrays e Slices: Homogêneos
// todos os elementos tem o mesmo tipo
// [1,2,3,4,5,6] - []int
// ["kaenna", "bento", "golang"] []string
// [1.1, 2.01, 3.1, 4.1, 5.1, 6.1] - []float

// 2 - Maps: Heterogêneos
// pode misturar tipos
// estrutura chave - valor
// [key] = value
// chave tem um tipo, e o valor pode ter outro
// map[string]int
// {"kaenna":28, "bento": 4}
// map[string]string
// {"kaenna": "gualoa", "bento": "gualoa"}

// Array

// Tamanho fixo, de zero ou mais elementos do mesmo tipo
// Acessamos os valores com índice: a[0], a[1]
// Função embutida len() retorna o tamanho do Array
// Por conta do tamanho fixo, não é tão usado. Só em casos específicos

// Slice

// Tipo o array, mas sem tamanho fixo
// Acessamos os valores com índice: a[0], a[1]
// Função embutida len() retorna o tamanho do Array
// Função append() usada para adicionar
