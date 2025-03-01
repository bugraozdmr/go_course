package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue : ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses : ")
	fmt.Scan(&revenue)

	fmt.Print("Tax Rate : ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	formattedMessage := fmt.Sprintf("Profit Ratio : %.4f", ratio)

	// fmt.Print("Profit : ", ratio)
	// fmt.Printf("Profit Ratio : %v", ratio)
	fmt.Printf(formattedMessage)

	// `ilede yazılabilir ve backtickler ozel karakterleri umursamazlar \n gibi`
}
