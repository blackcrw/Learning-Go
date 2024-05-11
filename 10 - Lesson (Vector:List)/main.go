package main

import "fmt"

// Criando, Adicionando e Lendo vetores/lista

func main() {
	var friends [5]string

	for i := 0; i < len(friends); i++ {
		fmt.Print("Digite o nome de um dos seus amigos: ")
		fmt.Scanf("%s", &friends[i])
	}

	fmt.Println("Seus amigos são: ")
	for _, friend := range friends {
		fmt.Printf("— %s\n", friend)
	}

	// arrayInicializado := [3]int{2, 4, 6}

	var matriz [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print("Digite um numero: ")
			fmt.Scanf("%d", &matriz[i][j])
		}
	}

	fmt.Println(matriz)

}
