package main

import (
	"errors"
	"fmt"
	"time"
)

func checkPower(hasPower bool) error {
	if !hasPower {
		return errors.New("No power, unable to cook!")
	}
	return nil
}

func addWater(water string) error {
	if water == "" {
		return errors.New("Please specify the amount of water!")
	}
	return nil
}

func checkItemAndTimer(itemToCook string, timer int) error {
	if itemToCook == "" {
		return errors.New("Please specify the item to cook!")
	}
	if itemToCook == "rice" && timer < 3 {
		return errors.New("Time to cook is insufficient")
	}
	return nil
}

func startCooking(itemToCook string, water string, timer int) error {
	if timer <= 0 {
		return errors.New("Timer must be greater than 0!")
	}

	fmt.Println("Power is available!")
	fmt.Printf("Adding %s cups of water...\n", water)
	fmt.Printf("Starting %s cooking...\n", itemToCook)
	// La fonction Sleep ne serait pas directement disponible en Go comme dans certains autres langages
	// La logique de pause peut être implémentée en utilisant time.Sleep ou d'autres méthodes asynchrones
	time.Sleep(time.Duration(timer) * time.Second) // En supposant que le minuteur est en secondes (converti en millisecondes)
	fmt.Printf("%s is cooked! Time to eat.\n", itemToCook)
	return nil
}

func cook(timer int, itemToCook string, water string, hasPower bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("An error occurred: %s\n", r)
		}
	}()

	if err := checkPower(hasPower); err != nil {
		panic(err.Error())
	}
	if err := addWater(water); err != nil {
		panic(err.Error())
	}
	if err := checkItemAndTimer(itemToCook, timer); err != nil {
		panic(err.Error())
	}
	if err := startCooking(itemToCook, water, timer); err != nil {
		panic(err.Error())
	}
}

func main() {
	cook(10, "rice", "", true)
	cook(10, "", "2c", true)
	cook(2, "tomato", "4c", true)
	cook(2, "rice", "4c", true)
	cook(10, "rice", "4c", true)
}
