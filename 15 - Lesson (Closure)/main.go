package main

import "fmt"

// Closures, in JavaScript this would be similar to an Arrow Function, and in Python it would be a lambda.

func main() {
	var num1, num2 int
	var operation string
	var methodOperation func(n1 int, n2 int) int // Here I declare a Closure.

	fmt.Print("Value 1: ")
	fmt.Scanf("%d", &num1)

	fmt.Print("Value 2: ")
	fmt.Scanf("%d", &num2)

	fmt.Print("Operation type (+, -, (x, *), /): ")
	fmt.Scanf("%s", &operation)

	if operation == "+" {
		methodOperation = func(n1 int, n2 int) int { // Here I create a variable that can store a pointer/function.
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
		fmt.Printf("Invalid operation '%s'.\n", operation)
	}

	ResultCalc := methodOperation(num1, num2) // Here I call the operation method, in this case methodOperation.

	fmt.Printf("—————————\n%d %s %d = %d\n", num1, operation, num2, ResultCalc)
}
