package main

import "fmt"

// Closures, no javascript isso seria similar uma Arrow Function e no python lambda.

func main() {
	var num1, num2 int
	var operation string
	var methodOperation func(n1 int, n2 int) int // Aqui eu instâncio uma Closure.

	fmt.Print("Valor 1: ")
	fmt.Scanf("%d", &num1)

	fmt.Print("Valor 2: ")
	fmt.Scanf("%d", &num2)

	fmt.Print("Tipo de operação (+, -, (x, *), /): ")
	fmt.Scanf("%s", &operation)

	if operation == "+" {
		methodOperation = func(n1 int, n2 int) int { // Aqui eu crio uma variável que é capaz de guardar um ponteiro/função nela.
			return n1 + n2
		}
	} else if operation == "-" {
		methodOperation = func(n1 int, n2 int) int {
			return n1 - n2
		}
	} else if operation == "x" || operation == "*" {
		methodOperation = func(n1 int, n2 int) int {
			return n1 * n2
		}
	} else if operation == "/" {
		methodOperation = func(n1 int, n2 int) int {
			return n1 / n2
		}
	} else {
		fmt.Printf("Operação '%s' inválida.\n", operation)
	}

	ResultCalc := methodOperation(num1, num2) // Aqui eu chamo o methodo operação, no caso methodOperation.

	fmt.Printf("—————————\n%d %s %d = %d\n", num1, operation, num2, ResultCalc)
}
