package main

import "fmt"

// Multiplication table using for loop

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &num)

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}
}
