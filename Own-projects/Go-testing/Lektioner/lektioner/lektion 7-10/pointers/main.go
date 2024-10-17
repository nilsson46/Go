package main

import (
	"fmt"
)

type Car struct {
	Brand     string
	ModelYear int
	Seats     int
}

func changeOneCar(cars []*Car) { // om jag hade lagt till []*Car efter () så hade jag ändrat originalet
	// Ändra originalet genom pekare
	cars[0].Brand = "Tesla"
	newCar := &Car{
		Brand:     "BMW",
		ModelYear: 2022,
		Seats:     5,
	}
	cars = append(cars, newCar) // Detta påverkar inte original-slicen utanför funktionen

	fmt.Println("Inside changeOneCar - After changes:")
	for _, car := range cars {
		fmt.Println(car.Brand)
	}
}

func main() {
	car1 := &Car{
		Brand:     "Toyota",
		ModelYear: 2010,
		Seats:     5,
	}
	car2 := &Car{
		Brand:     "Honda",
		ModelYear: 2015,
		Seats:     5,
	}
	car3 := &Car{
		Brand:     "Ford",
		ModelYear: 2018,
		Seats:     5,
	}

	var cars = []*Car{car1, car2, car3} // Original-slice med pekare

	fmt.Println("Original slice - Before changeOneCar:")
	for _, car := range cars {
		fmt.Println(car.Brand)
	}

	changeOneCar(cars) // Ändrar namnet på den första bilen i original-slicen

	fmt.Println("Original slice - After changeOneCar:")
	for _, car := range cars {
		fmt.Println(car.Brand)
	}

	newCar := &Car{
		Brand:     "Audi",
		ModelYear: 2021,
		Seats:     5,
	}

	cars = append(cars, newCar) // Lägger till en ny bil i original-slicen

	fmt.Println("Original slice - After appending newCar:")
	for _, car := range cars {
		fmt.Println(car.Brand)
	}
}
