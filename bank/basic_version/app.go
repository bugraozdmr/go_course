package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Tosuncuk Bank")

	for {
		fmt.Println("\nWhat do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				return
			}
			accountBalance += depositAmount
			fmt.Println("Deposit successful. Your new balance is:", accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter withdraw amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds.")
			} else {
				accountBalance -= withdrawAmount
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
