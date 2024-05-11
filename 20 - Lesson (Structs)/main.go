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
		return errors.New("the car's speed is already at its limit") // Creating an error message.
	}
}

func (car *Car) Brake() error {
	if car.velocity > 0 {
		car.velocity -= 5
		return nil
	} else {
		return errors.New("the car is already stopped")
	}
}

func main() {
	car := Car{model: "Palio", brand: "Fiat"}
	option := 0

	for option != 3 {
		fmt.Println("What do you want to do?")
		fmt.Println("1 - Speed up")
		fmt.Println("2 - Brake")
		fmt.Println("3 - Exit")
		fmt.Scanf("%d", &option)
		var err error = nil

		switch option {
		case 1:
			err = car.SpeedUp()
		case 2:
			err = car.Brake()
		}

		if err != nil {
			fmt.Printf("** Unable to perform this action: %s ** \n", err)
		} else if option != 3 {
			fmt.Printf("The %s car from %s is going at %.2f KM/h \n", car.model, car.brand, car.velocity)
		}
	}
}
