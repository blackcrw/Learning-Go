package main

import "fmt"

// Função

func main() {
	var num1 int
	var num2 int
	var operation string

	fmt.Print("Valor 1: ")
	fmt.Scanf("%d", &num1)

	fmt.Print("Valor 2: ")
	fmt.Scanf("%d", &num2)

	fmt.Print("Tipo de operação (+, -, (x, *), /): ")
	fmt.Scanf("%s", &operation)

	if operation == "+" {
		add(num1, num2)
	} else if operation == "-" {
		// resultCalc := subtract(num1, num2)
		// fmt.Printf("Resultado: %d - %d = %d\n", num1, num2, resultCalc)

		fmt.Printf("Resultado: %d - %d = %d\n", num1, num2, subtract(num1, num2))
	} else if operation == "*" || operation == "x" {
		fmt.Printf("Resultado: %d x %d = %d\n", num1, num2, multiply(num1, num2))
	} else if operation == "/" {
		resultCalc, resultRest := division(num1, num2)
		fmt.Printf("Resultado: %d / %d = %d; Resto: %d\n", num1, num2, resultCalc, resultRest)
	} else {
		fmt.Println("Operação errada!")
	}

}

func add(n1 int, n2 int) { // Isto é uma função, mas está função não retorna nenhum valor. Ela apenas printa.
	fmt.Printf("Resultado: %d + %d = %d\n", n1, n2, n1+n2)
}

func subtract(n1 int, n2 int) int { // Já está função ela tem um retorno. Para haver algum retorno é preciso informar qual o tipo do retorno, se é int, string, boolean e etc...
	return n1 + n2
}

func multiply(n1 int, n2 int) (resultCalc int) { // Aqui estou setando a variável 'resultCalc' no instanciamento da função.
	resultCalc = n1 * n2
	return resultCalc
}

func division(n1 int, n2 int) (int, int) { // Aqui eu informo que terá dois valores a serem retornados e os dois será inteiros.
	resultCalc := n1 / n2
	resultRest := n1 % n2

	return resultCalc, resultRest
}
