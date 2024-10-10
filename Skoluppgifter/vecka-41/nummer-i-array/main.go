package main

import "fmt"

func addNumberFromInput() int {
	var number int
	fmt.Print("Enter a number: ")

	fmt.Scan(&number)
	if number < 0 {
		fmt.Println("The number must be positive.")
		return addNumberFromInput()
	}
	return number
}

func main() {
	var numberArray [4]int
	for i := 0; i < 4; i++ {
		number := addNumberFromInput()
		numberArray[i] = number
	}
	var maxNumber int
	for _, number := range numberArray {
		if number > maxNumber {
			maxNumber = number
		}
	}
	fmt.Println("The largest number is:", maxNumber)
}
