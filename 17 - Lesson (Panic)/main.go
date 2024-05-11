package main

import "fmt"

func main() {
	var num int

	fmt.Println("Olá seja bem-vindo!")
	defer fmt.Println("Código chegou ao fim! Obrigado.")

	fmt.Print("Digite um numero acima de 5: ")
	fmt.Scanf("%d", &num)

	if num <= 5 {
		panic("Ocorreu um erro! Número inválido.")
	} else {
		fmt.Println("Hmmm! ótimo numero.")
	}
}
