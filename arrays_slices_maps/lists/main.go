package main

import "fmt"

func main() {
	prices := []float64{10.99,8.99}
	fmt.Println(prices[0:1])
	prices[1] = 12.22

	prices = append(prices,2.99)

	fmt.Println(prices)
}

/*
type Product struct {
	title string
	id string
	price float64
}

func main() {
	var productNames []string

	prices := [4]float64{10.99,20.22}
	fmt.Println(prices)
	fmt.Println(productNames)

	// featuredPrices := prices[1:2]
	
	// sona kadar 1'den
	featuredPrices := prices[1:]
	fmt.Println(featuredPrices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))
}
	*/