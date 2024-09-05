package main

import "fmt"

//loops
//laços de repetição
//Repetir tarefas

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println("Sum:", sum)  //Mostrando como o for funciona
		fmt.Println("Indice:", i) //Mostrando como o for funciona
		sum += i
		//é a mesma coisa que: sum = sum + i
		// sum -= i -> sum = sum -i

		// É como se ao final do loop
		//fosse adicionado 1 ao indice i
		// i++
		// i = i + 1
	}

	fmt.Println(sum)

	//Loop infinito
	// for {
	// 	fmt.Println("Golang do Zero")
	// 	time.Sleep(2 * time.Second)
	// }

	// for range
	// frutas := []string{"laranja", "maça", "banana", "uva", "kiwi"}
	// for _, fruta := range frutas { // _, significa a posição
	// 	fmt.Println(fruta)
	// }
}

//numero := 100 (Variavel criada dentro do for só existe dentro dele)
//fmt.Println("Numero:", numero) -- (Caso eu queira printar fora, ele dará erro)
