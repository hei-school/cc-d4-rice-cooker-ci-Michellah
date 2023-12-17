package main

import (
	"errors"
	"fmt"
	"time"
)

func CheckPower(hasPower bool) error {
	if !hasPower {
		return errors.New("No power, unable to cook!")
	}
	return nil
}

func AddWater(water string) error {
	if water == "" {
		return errors.New("Please specify the amount of water!")
	}
	return nil
}

func CheckItemAndTimer(itemToCook string, timer int) error {
	if itemToCook == "" {
		return errors.New("Please specify the item to cook!")
	}
	if itemToCook == "rice" && timer < 3 {
		return errors.New("Time to cook is insufficient")
	}
	return nil
}

func StartCooking(itemToCook string, water string, timer int) error {
	if timer <= 0 {
		return errors.New("Timer must be greater than 0!")
	}

	fmt.Println("Power is available!")
	fmt.Printf("Adding %s cups of water...\n", water)
	fmt.Printf("Starting %s cooking...\n", itemToCook)
	time.Sleep(time.Duration(timer) * time.Second) 
	fmt.Printf("%s is cooked! Time to eat.\n", itemToCook)
	return nil
}

func Cook(timer int, itemToCook string, water string, hasPower bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("An error occurred: %s\n", r)
		}
	}()

	if err := CheckPower(hasPower); err != nil {
		panic(err.Error())
	}
	if err := AddWater(water); err != nil {
		panic(err.Error())
	}
	if err := CheckItemAndTimer(itemToCook, timer); err != nil {
		panic(err.Error())
	}
	if err := StartCooking(itemToCook, water, timer); err != nil {
		panic(err.Error())
	}
}

func main() {
	Cook(10, "rice", "", true)
	Cook(10, "", "2c", true)
	Cook(2, "tomato", "4c", true)
	Cook(2, "rice", "4c", true)
	Cook(10, "rice", "4c", true)
}
