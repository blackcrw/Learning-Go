package main

import "fmt"

// defer is similar to Java's finally.

func main() {
	fmt.Println("Opening database connection.")
	// This line of code will be deferred to the end of this process. In other words, it will only be executed at the end of the function.
	defer fmt.Println("Closing database connection.")
	execQuery()
}

func execQuery() {
	fmt.Println("Executing database query!")
}
