package main

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	moreNumbers := []int{5, 1, 2, 3}

	// address değerini verdik sürekli sürekli üretmesin diyerek
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println("doubled numbers :", doubled)
	fmt.Println("tripled numbers :", tripled)

	transformerFn1 := getTrasnformerFunction(&numbers)
	transformerFn2 := getTrasnformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println("transformed numbers :", transformedNumbers)
	fmt.Println("more t numbers :", moreTransformedNumbers)

	// anonymous function has to match the exact same number of arguments in function
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 13
	})

	fmt.Println(transformed)

	//
	doublee := createTransformer(2)
	triplee := createTransformer(3)

	fmt.Println("check out", transformNumbers(&numbers, doublee))
	fmt.Println("check out", transformNumbers(&numbers, triplee))

	//
	fact := factorial(4)
	fmt.Println("Factorial: ", fact)

}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func getTrasnformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {

		return double
	} else { // go böyle istermiş parantezler aynı satırda olcak
		return triple
	}

	return double
}

func createTransformer(factor int) func(int) int {
	// outer scope can be used in inner scope
	return func(number int) int {
		return number * factor
	}
}

func factorial(number int) int {
	/*result := 1

	for i := 1; i <= number; i++ {
		result = result * i
	}

	return result*/
	// recuursion

	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}
