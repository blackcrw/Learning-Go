package main

import "fmt"

// Criando vetor/slice/lista, adicionando valores num vetor/lista com append.

func main() {
	var name string
	friends := make([]string, 3)

	for i := 0; name != "q"; i++ {
		fmt.Print("Digite o nome de um amigo e digite 'q' para parar: ")
		fmt.Scanf("%s", &name)
		if name != "q" {
			if i < 3 {
				friends[i] = name
			} else {
				// friends = append(friends, name, "João", "Maria", "Mario")
				friends = append(friends, name)
			}
		}
	}

	fmt.Printf("\nVocê tem %v amigo(s) \n", len(friends))
	for _, friend := range friends {
		fmt.Printf("— %s \n", friend)
	}

	// other_type_slice := friends[:3] // Quando nenhum valor é colocado antes do ':' quer dizer que é sempre 0

	other_type_slice := friends[1:3]
	fmt.Println(other_type_slice)
}
