package main

import "fmt"

// make kaç item olacağını biliyorsan daha iyi diyor

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}


func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Grant")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.6
	courseRatings["laravel"] = 4.2
	courseRatings["react"] = 4.1

	// fmt.Println(courseRatings)
	courseRatings.output()

	for index, value := range userNames{
		fmt.Println("index : ",index)
		fmt.Println("value : ",value)
	}
}