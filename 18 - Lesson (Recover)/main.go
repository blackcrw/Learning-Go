package main

import "fmt"

// Recover
// Error handling.

func main() {
	var num int

	fmt.Println("Welcome!")
	// Here I create an Anonymous Function (Closure)
	defer func() {
		// I create a variable called 'ex' and assign the value that recover returns to it. If it's not nil, it executes.
		if ex := recover(); ex != nil {
			fmt.Printf("\nAn error occurred: %s\n", ex)
		} else {
			fmt.Println("No errors occurred!")
		}
	}()

	fmt.Print("Enter a number greater than 5: ")
	fmt.Scanf("%d", &num)

	if num <= 5 {
		panic("Invalid number!") // Here I create the error message that will appear when recover is executed
	} else {
		fmt.Println("Nice number!")
	}

}
