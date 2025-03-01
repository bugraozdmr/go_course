package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5

	// bu şekil tutatabilirsin
	var insvestmentAmount float64
	// type'i go tarafından belirlensin
	expectedReturnRate := 5.5
	var years float64

	// var a, b float64 = 10, 20
	// var c, d = 10, "20"

	// scan -- pointer ile beraber kullanmalısın
	fmt.Print("Investment Amount : ")
	fmt.Scan(&insvestmentAmount)

	fmt.Print("Investment Years : ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate : ")
	fmt.Scan(&expectedReturnRate)

	// var futureValue = float64(insvestmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	var futureValue = insvestmentAmount * math.Pow(1+expectedReturnRate/100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
