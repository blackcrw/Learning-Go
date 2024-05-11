package main

import "fmt"

// Recover
// Tratamento de erros.

func main() {
	var num int

	fmt.Println("Seja bem vindo!")
	// Aqui eu crio uma Função Anonima (Closure)
	defer func() {
		// Eu crio uma variável chamada 'ex' e coloco o valor que o recover tiver dentro dela. Caso não seja nil ele executa.
		if ex := recover(); ex != nil {
			fmt.Printf("\nOcorreu um erro: %s\n", ex)
		} else {
			fmt.Println("Nenhum erro ocorreu!")
		}
	}()

	fmt.Print("Digite um numero maior que 5: ")
	fmt.Scanf("%d", &num)

	if num <= 5 {
		panic("Número inválido!") // Aqui eu crio a mensagem de erro que parecerá quando o recover for executado
	} else {
		fmt.Println("Belo numero!")
	}

}
