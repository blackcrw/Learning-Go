package main

import "fmt"

// Utilizando switch e case

func main() {
	var mes int

	fmt.Print("Digite o numero de um mês: ")
	fmt.Scanf("%d", &mes)

	switch mes {
	case 1:
		fmt.Println("\nEste é o mês de Janeiro")
	case 2:
		fmt.Println("\nEste é o mês de Fevereiro")
	case 3:
		fmt.Println("\nEste é o mês de Março")
	case 4:
		fmt.Println("\nEste é o mês de Abril")
	case 5:
		fmt.Println("\nEste é o mês de Maio")
	case 6:
		fmt.Println("\nEste é o mês de Junho")
	case 7:
		fmt.Println("\nEste é o mês de Julho")
	case 8:
		fmt.Println("\nEste é o mês de Agosto")
	case 9:
		fmt.Println("\nEste é o mês de Setembro")
	case 10:
		fmt.Println("\nEste é o mês de Outubro")
	case 11:
		fmt.Println("\nEste é o mês de Novembro")
	case 12:
		fmt.Println("\nEste é o mês de Dezembro")
	default:
		fmt.Println("\nEste mês é inválido")
	}

	switch mes {
	case 1, 2, 3:
		fmt.Println("Primeiro trimestre")
	case 4, 5, 6:
		fmt.Println("Segundo trimestre")
	case 7, 8, 9:
		fmt.Println("Terceiro trimestre")
	case 10, 11, 12:
		fmt.Println("Quarto trimestre")
	}

}
