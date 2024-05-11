package main

import "fmt"

// Using switch and case

func main() {
	var month int

	fmt.Print("Enter the number of a month: ")
	fmt.Scanf("%d", &month)

	switch month {
	case 1:
		fmt.Println("\nThis is January")
	case 2:
		fmt.Println("\nThis is February")
	case 3:
		fmt.Println("\nThis is March")
	case 4:
		fmt.Println("\nThis is April")
	case 5:
		fmt.Println("\nThis is May")
	case 6:
		fmt.Println("\nThis is June")
	case 7:
		fmt.Println("\nThis is July")
	case 8:
		fmt.Println("\nThis is August")
	case 9:
		fmt.Println("\nThis is September")
	case 10:
		fmt.Println("\nThis is October")
	case 11:
		fmt.Println("\nThis is November")
	case 12:
		fmt.Println("\nThis is December")
	default:
		fmt.Println("\nThis month is invalid")
	}

	switch month {
	case 1, 2, 3:
		fmt.Println("First quarter")
	case 4, 5, 6:
		fmt.Println("Second quarter")
	case 7, 8, 9:
		fmt.Println("Third quarter")
	case 10, 11, 12:
		fmt.Println("Fourth quarter")
	}

}
