package utils

import "fmt"

// Para yatırma işlemi
func DepositMoney(accountBalance float64, filename string) float64 {
	var depositAmount float64
	fmt.Print("Enter deposit amount: ")
	fmt.Scan(&depositAmount)

	if depositAmount <= 0 {
		fmt.Println("Invalid amount. Must be greater than 0")
		return accountBalance
	}

	accountBalance += depositAmount
	writeBalanceToFile(filename, accountBalance)
	fmt.Println("Deposit successful. Your new balance is:", accountBalance)
	return accountBalance
}

// Para çekme işlemi
func WithdrawMoney(accountBalance float64, filename string) float64 {
	var withdrawAmount float64
	fmt.Print("Enter withdraw amount: ")
	fmt.Scan(&withdrawAmount)

	if withdrawAmount > accountBalance {
		fmt.Println("Insufficient funds.")
		return accountBalance
	}

	accountBalance -= withdrawAmount
	writeBalanceToFile(filename, accountBalance)
	fmt.Println("Withdrawal successful. Your new balance is:", accountBalance)
	return accountBalance
}
