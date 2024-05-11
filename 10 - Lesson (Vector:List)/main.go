package main

import "fmt"

// Creating, Adding, and Reading Arrays/Lists

func main() {
	var friends [5]string

	for i := 0; i < len(friends); i++ {
		fmt.Print("Enter the name of one of your friends: ")
		fmt.Scanf("%s", &friends[i])
	}

	fmt.Println("Your friends are: ")
	for _, friend := range friends {
		fmt.Printf("— %s\n", friend)
	}

	// initializedArray := [3]int{2, 4, 6}

	var matrix [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print("Enter a number: ")
			fmt.Scanf("%d", &matrix[i][j])
		}
	}

	fmt.Println(matrix)
}
