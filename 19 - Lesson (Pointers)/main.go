package main

import "fmt"

// Pointers

func main() {
	var balance float64

	fmt.Print("Enter your balance: ")
	fmt.Scanf("%f", &balance)

	calcInterest(&balance) // Here I'm not passing the variable, but its pointer in memory.

	fmt.Printf("Your balance with interest is: $ %.2f\n", balance)
}

func calcInterest(balance *float64) {
	*balance += *balance * 0.02 // Here I'm not changing the value of the variable, but the value in memory.
}
