package main

import "fmt"

type Product struct {
	name string
	price float64
}

func main() {
	// 1
	hobbies := [3]string{"Sports", "Cooking", "Reading"}
	fmt.Println(hobbies)

	// 2
	fmt.Println(hobbies[1])
	fmt.Println(hobbies[1:])

	// 3
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	// 4
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:]
	fmt.Println(mainHobbies)

	// 5
	courseGoals := []string{"Learn Go", "Learn Russian"}
	fmt.Println(courseGoals)

	// 6
	courseGoals[1] = "Learn with us"
	courseGoals = append(courseGoals, "Learn All the basics")
	fmt.Println(courseGoals)

	// 7
	products := []Product{
		{"MacBook", 40000},
		{"iPad", 17000},
	}
	fmt.Println(products)

	newProduct := Product{
		"iPhone 13",
		33000,
	}
	products = append(products, newProduct)
	fmt.Println(products)

	// 8
	prices := []float64{10.99,12.2}
	discountedPrices := []float64{101.2,11.3}
	// ... ile seperate ediyor
	prices = append(prices, discountedPrices...)

	fmt.Println(prices)
}