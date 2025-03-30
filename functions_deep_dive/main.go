package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 45, 6}
	sum := sumup("Hello yall", 1, 2, 3, 4, 56, 2)

	fmt.Println(sum)

	// dagitmak
	sumupme := sumup("Another One", numbers...)
	fmt.Println(sumupme)
}

// variadic functions
func sumup(message string, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
