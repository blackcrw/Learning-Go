package main

import "fmt"

/*
	In GoLang, we don't have an explicit while loop.
	The while loop has been replaced by the for loop itself! In other words, the for loop also serves as a while loop.
	Below is an example of using the for loop as a while loop.
*/

func main() {
	/*
		i := 0
		for true {
			fmt.Printf("Count: %d\n", i)
			i++
		}
	*/

	for i := 0; true; i++ {
		fmt.Printf("Count: %d\n", i)
	}
}
