package main

import "fmt"

// defer é parecido com finally do Java.

func main() {
	fmt.Println("Abrindo conexão com banco de dados.")
	// Essa linha de código será jogada para o final desse processo. Ou seja: Só será executada no final da função.
	defer fmt.Println("Fechando a conexão com banco de dados.")
	execQuery()
}

func execQuery() {
	fmt.Println("Executando busca no banco de dados!")
}
