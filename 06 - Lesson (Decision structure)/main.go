package main

import "fmt"

// Utilizando if, else if e else

func main() {
	/*
		var idade int

		fmt.Print("Digite sua idade: ")
		fmt.Scanf("%d", &idade)

		if idade >= 18 {
			fmt.Println("\nVocê é maior de idade")
		} else {
			fmt.Println("\nVocê e menor de idade")
		}
	*/

	var num int

	fmt.Print("Digite um numero: ")
	fmt.Scanf("%d", &num) // '%d' informa para o Scanf que o que estamos estaneando é um inteiro.

	if num < 0 {
		fmt.Println("Este é um valor negativo.")
	} else if num >= 0 && num <= 9 { // '&&' seria a mesma coisa que 'and' no Python.
		fmt.Println("Este é um valor positivo de apenas de um dígito.")
	} else {
		fmt.Println("Este é um valor positivo e contém mais de um dígito.")
	}

}
