package main

import (
	"fmt"
	"os"
)

var nameSlice []string

func addNameFromInput() string {
	var descision string
	var name string
	fmt.Print("Do you want to add a name? (y/n): ")
	fmt.Scan(&descision)
	if descision == "n" {
		fmt.Print(nameSlice)
		os.Exit(0)
	} else if descision != "y" {
		fmt.Print("Enter a name: ")
		fmt.Scan(&name)
		if name == "" {
			fmt.Println("The name must not be empty.")
			return addNameFromInput()
		}
		return name
	}

	fmt.Print("Enter a name: ")

	fmt.Scan(&name)
	if name == "" {
		fmt.Println("The name must not be empty.")
		return addNameFromInput()
	}
	return name
}

func main() {

	for i := 0; i < 4; i++ {
		name := addNameFromInput()
		nameSlice = append(nameSlice, name)
	}
	fmt.Println("The names are:", nameSlice)
}
