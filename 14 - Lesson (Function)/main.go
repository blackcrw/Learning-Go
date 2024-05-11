package main

import "fmt"

// Function

func main() {
	var num1 int
	var num2 int
	var operation string

	fmt.Print("Value 1: ")
	fmt.Scanf("%d", &num1)

	fmt.Print("Value 2: ")
	fmt.Scanf("%d", &num2)

	fmt.Print("Operation type (+, -, (x, *), /): ")
	fmt.Scanf("%s", &operation)

	if operation == "+" {
		add(num1, num2)
	} else if operation == "-" {
		// resultCalc := subtract(num1, num2)
		// fmt.Printf("Result: %d - %d = %d\n", num1, num2, resultCalc)

		fmt.Printf("Result: %d - %d = %d\n", num1, num2, subtract(num1, num2))
	} else if operation == "*" || operation == "x" {
		fmt.Printf("Result: %d x %d = %d\n", num1, num2, multiply(num1, num2))
	} else if operation == "/" {
		resultCalc, resultRest := division(num1, num2)
		fmt.Printf("Result: %d / %d = %d; Remainder: %d\n", num1, num2, resultCalc, resultRest)
	} else {
		fmt.Println("Wrong operation!")
	}

}

func add(n1 int, n2 int) { // This is a function, but this function does not return any value. It only prints.
	fmt.Printf("Result: %d + %d = %d\n", n1, n2, n1+n2)
}

func subtract(n1 int, n2 int) int { // This function has a return value. To have a return value, it is necessary to inform the return type, whether it is int, string, boolean, etc...
	return n1 + n2
}

func multiply(n1 int, n2 int) (resultCalc int) { // Here I am setting the variable 'resultCalc' during the function instantiation.
	resultCalc = n1 * n2
	return resultCalc
}

func division(n1 int, n2 int) (int, int) { // Here I inform that there will be two values to be returned and both will be integers.
	resultCalc := n1 / n2
	resultRest := n1 % n2

	return resultCalc, resultRest
}
