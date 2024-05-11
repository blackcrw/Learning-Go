package main

import (
	"fmt"
	"strconv"
)

func main() {
	var text string

	fmt.Print("Digite o primeiro numero: ")
	fmt.Scanf("%s", &text)

	// num, _ := strconv.Atoi(text)
	num, _ := strconv.ParseInt(text, 10, 64) // Utilizo a LIB strconv para fazer a conversão da string para um int de 64 bits

	fmt.Printf("Numero Convertido: %v\n", num)
}
