package main

import "fmt"

// Using if, else if, and else

func main() {
	/*
		var age int

		fmt.Print("Enter your age: ")
		fmt.Scanf("%d", &age)

		if age >= 18 {
			fmt.Println("\nYou are an adult")
		} else {
			fmt.Println("\nYou are a minor")
		}
	*/

	var num int

	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &num) // '%d' tells Scanf that we are scanning an integer.

	if num < 0 {
		fmt.Println("This is a negative value.")
	} else if num >= 0 && num <= 9 { // '&&' is equivalent to 'and' in Python.
		fmt.Println("This is a positive value with only one digit.")
	} else {
		fmt.Println("This is a positive value with more than one digit.")
	}

}
