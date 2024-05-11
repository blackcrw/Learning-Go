package main

import (
	"container/list"
	"fmt"
)

// Isso é uma lista ligada

func main() {
	listNum := list.New()     // Instancio a lista
	el := listNum.PushBack(1) // Adiciono o valor 1 no final da lista
	listNum.PushFront(0)      // Adiciono o valor 0 no inicio da lista
	listNum.PushBack(2)       // Adiciono o valor 2 no final da lista

	// percorro a lista do inicio ao fim. nil é a mesma coisa que null, ou no python None. O for só para de percorrer a lista quando ela retorna um nil (null/none).
	for e := listNum.Front(); e != nil; e.Next() {
		fmt.Printf("Numero da lista: %d", e.Value) // Estou printando os valores que estão na lista.
	}

	listNum.Remove(el) // Estou removendo o valor 1 da lista

	for e := listNum.Front(); e != nil; e.Next() {
		fmt.Printf("Numero da lista: %d", e.Value)
	}
}
