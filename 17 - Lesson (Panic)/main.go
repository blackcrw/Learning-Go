package main

import "fmt"

func main() {
	var num int

	fmt.Println("Hello, welcome!")
	defer fmt.Println("Code has reached the end! Thank you.")

	fmt.Print("Enter a number above 5: ")
	fmt.Scanf("%d", &num)

	if num <= 5 {
		panic("An error occurred! Invalid number.")
	} else {
		fmt.Println("Hmm! Great number.")
	}
}
