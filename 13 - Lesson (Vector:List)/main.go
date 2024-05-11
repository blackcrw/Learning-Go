package main

import (
	"container/list"
	"fmt"
)

// This is a linked list

func main() {
	listNum := list.New()     // Instantiate the list
	el := listNum.PushBack(1) // Add the value 1 to the end of the list
	listNum.PushFront(0)      // Add the value 0 to the beginning of the list
	listNum.PushBack(2)       // Add the value 2 to the end of the list

	// Traverse the list from start to end. nil is equivalent to null, or None in Python. The loop stops when the list returns nil (null/none).
	for e := listNum.Front(); e != nil; e = e.Next() {
		fmt.Printf("Number from the list: %d", e.Value) // Print the values in the list.
	}

	listNum.Remove(el) // Remove the value 1 from the list

	for e := listNum.Front(); e != nil; e = e.Next() {
		fmt.Printf("Number from the list: %d", e.Value)
	}
}
