package main

import "fmt"

/*
	Na GoLang não temos um while explicito.
	O while foi substituido pelo próprio for! Ou seja: O for também faz a função de um While.
	Abaixo deixarei um exemplo de uso do for para while.
*/

func main() {
	/*
		i := 0
		for true {
			fmt.Printf("Contagem: %d\n", i)
			i++
		}
	*/

	for i := 0; true; i++ {
		fmt.Printf("Contagem: %d\n", i)
	}
}
