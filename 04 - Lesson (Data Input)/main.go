package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Print("What is the first number? ")
	fmt.Scanln(&num1)

	fmt.Print("What is the second number? ")
	fmt.Scanln(&num2)

	fmt.Printf("\n%d + %d = %d\n", num1, num2, num1+num2)
	fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	fmt.Printf("%d x %d = %d\n", num1, num2, num1*num2)
	fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
	fmt.Printf("%d %% %d = %d\n", num1, num2, num1%num2)

	// increment := num1
	// increment += num2
	increment := num1 + num2
	fmt.Printf("\nThe increment of %d with %d is %d\n", num1, num2, increment)

	// decrement := num1
	// decrement -= num2
	decrement := num1 - num2
	fmt.Printf("The decrement of %d with %d is %d\n", num1, num2, decrement)
}
