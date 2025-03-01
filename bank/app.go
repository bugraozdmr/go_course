package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	const filename = "balance.txt"
	accountBalance := readBalanceFromFile(filename)

	fmt.Println("Welcome to Tosuncuk Bank")

	for {
		choice := showMenuAndGetChoice()

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)

		case 2:
			var depositAmount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
			} else {
				accountBalance += depositAmount
				writeBalanceToFile(filename, accountBalance)
				fmt.Println("Deposit successful. Your new balance is:", accountBalance)
			}

		case 3:
			var withdrawAmount float64
			fmt.Print("Enter withdraw amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds.")
			} else {
				accountBalance -= withdrawAmount
				writeBalanceToFile(filename, accountBalance)
				fmt.Println("Withdrawal successful. Your new balance is:", accountBalance)
			}

		case 4:
			fmt.Println("Thank you for using Tosuncuk Bank. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please choose a valid option.")
		}
	}
}

func showMenuAndGetChoice() int {
	var choice int
	fmt.Println("\nWhat do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)
	return choice
}

func readBalanceFromFile(filename string) float64 {
	var accountBalance float64

	// Open the file for reading and writing, creating the file if it doesn't exist
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err) // Log any error while opening the file and stop the program
	}
	defer file.Close() // Ensure the file is closed after function execution

	// Read the balance from the file
	_, err = fmt.Fscanf(file, "%f", &accountBalance)
	if err != nil {
		// If there's an error reading (e.g., the file is empty), set the balance to 1000
		accountBalance = 1000
	}
	return accountBalance
}

func writeBalanceToFile(filename string, balance float64) {
	// Open the file for reading and writing, truncating the file to zero if it already exists
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err) // Log any error while opening the file and stop the program
	}
	defer file.Close() // Ensure the file is closed after function execution

	// Write the balance to the file -- sıfırdan yazıyor gibi append degil yani
	_, err = file.WriteString(fmt.Sprintf("%f", balance))
	if err != nil {
		log.Fatal(err) // Log any error while writing and stop the program
	}
}
