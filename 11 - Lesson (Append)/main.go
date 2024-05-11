package main

import "fmt"

// Creating a vector/slice/list, adding values to a vector/list using append.

func main() {
	var name string
	friends := make([]string, 3)

	for i := 0; name != "q"; i++ {
		fmt.Print("Enter the name of a friend and enter 'q' to stop: ")
		fmt.Scanf("%s", &name)
		if name != "q" {
			if i < 3 {
				friends[i] = name
			} else {
				// friends = append(friends, name, "John", "Maria", "Mario")
				friends = append(friends, name)
			}
		}
	}

	fmt.Printf("\nYou have %v friend(s) \n", len(friends))
	for _, friend := range friends {
		fmt.Printf("— %s \n", friend)
	}

	// other_type_slice := friends[:3] // When no value is placed before ':' it means it's always 0

	other_type_slice := friends[1:3]
	fmt.Println(other_type_slice)
}
