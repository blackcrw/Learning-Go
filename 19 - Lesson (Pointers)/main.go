package main

import "fmt"

// Ponteiros

func main() {
	var saldo float64

	fmt.Print("Digite seu saldo: ")
	fmt.Scanf("%f", &saldo)

	calcRend(&saldo) // Aqui eu não estou informando a variável, e sim o ponteiro dela na memória.

	fmt.Printf("Seu saldo com rendimento é : R$ %.2f\n", saldo)
}

func calcRend(saldo *float64) {
	*saldo += *saldo * 0.02 // Aqui não estou alterando o valor da variável, e sim o valor dela na memória.
}
