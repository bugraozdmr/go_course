package main

import (
	"fmt"

    "my-bank/utils"
)

func main() {
	const filename = "balance.txt"
	accountBalance := utils.ReadBalanceFromFile(filename)

	fmt.Println("Welcome to Tosuncuk Bank")

	for {
		choice := utils.ShowMenuAndGetChoice()

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)

		case 2:
			accountBalance = utils.DepositMoney(accountBalance, filename)

		case 3:
			accountBalance = utils.WithdrawMoney(accountBalance, filename)

		case 4:
			fmt.Println("Thank you for using Tosuncuk Bank. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please choose a valid option.")
		}
	}
}
