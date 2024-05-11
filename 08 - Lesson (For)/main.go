package main

import "fmt"

// Tabuada com for

func main() {
	var num int

	fmt.Print("Digite um numero: ")
	fmt.Scanf("%d", &num)

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}
}
