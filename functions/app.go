package main

import "fmt"

func getInput(prompt string) float64 {
	var input float64
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}

func calculate(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func main() {
	revenue := getInput("Revenue: ")
	expenses := getInput("Expenses: ")
	taxRate := getInput("Tax Rate: ")

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)

	formattedMessage := fmt.Sprintf("Ebt : %.4f\nProfit : %.4f\nProfit Ratio: %.4f", ebt, profit, ratio)
	fmt.Println(formattedMessage)
}
