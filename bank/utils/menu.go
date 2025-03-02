package utils

import "fmt"
import "github.com/Pallinder/go-randomdata"

func ShowMenuAndGetChoice() int {
	var choice int
	fmt.Println("Here is the silly name : ", randomdata.SillyName())
	fmt.Println("\nWhat do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)
	return choice
}
