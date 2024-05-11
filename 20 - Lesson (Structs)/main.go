package main

import (
	"errors"
	"fmt"
)

// Structs

type Car struct {
	model    string
	brand    string
	velocity float32
}

func (car *Car) SpeedUp() error {
	if car.velocity < 20 {
		car.velocity += 5
		return nil
	} else {
		return errors.New("a velocidade do carro já está no seu limite") // Criando uma mensagem de erro.
	}
}

func (car *Car) Brake() error {
	if car.velocity > 0 {
		car.velocity -= 5
		return nil
	} else {
		return errors.New("o carro já está parado")
	}
}

func main() {
	car := Car{model: "Palio", brand: "Fiat"}
	option := 0

	for option != 3 {
		fmt.Println("O que desejas fazer ?!")
		fmt.Println("1 - Acelerar")
		fmt.Println("2 - Freiar")
		fmt.Println("3 - Sair")
		fmt.Scanf("%d", &option)
		var err error = nil

		switch option {
		case 1:
			err = car.SpeedUp()
		case 2:
			err = car.Brake()
		}

		if err != nil {
			fmt.Printf("** Não foi possível executar essa ação: %s ** \n", err)
		} else if option != 3 {
			fmt.Printf("O carro %s da marca %s está a %.2f KM/h \n", car.model, car.brand, car.velocity)
		}
	}
}
