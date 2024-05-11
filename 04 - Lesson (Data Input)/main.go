package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Print("Qual o primeiro numero ? ")
	fmt.Scanln(&num1)

	fmt.Print("Qual o segundo numero ? ")
	fmt.Scanln(&num2)

	fmt.Printf("\n%d + %d = %d\n", num1, num2, num1+num2)
	fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	fmt.Printf("%d x %d = %d\n", num1, num2, num1*num2)
	fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
	fmt.Printf("%d %% %d = %d\n", num1, num2, num1%num2)

	// incremento := num1
	// incremento += num2
	incremento := num1 + num2
	fmt.Printf("\nO incremento de %d com %d é %d\n", num1, num2, incremento)

	// decremento := num1
	// decremento -= num2
	decremento := num1 - num2
	fmt.Printf("O decremento de %d com %d é %d\n", num1, num2, decremento)
}
