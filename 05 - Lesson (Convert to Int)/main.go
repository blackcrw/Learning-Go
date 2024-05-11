package main

import (
	"fmt"
	"strconv"
)

func main() {
	var text string

	fmt.Print("Enter the first number: ")
	fmt.Scanf("%s", &text)

	// num, _ := strconv.Atoi(text)
	num, _ := strconv.ParseInt(text, 10, 64) // Using the strconv library to convert the string to a 64-bit integer

	fmt.Printf("Converted Number: %v\n", num)
}
