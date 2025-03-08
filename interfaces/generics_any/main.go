package main

import "fmt"

func PrintValue(value any) {
	fmt.Println("Değer:", value)
}

func main() {
	PrintValue(42)        // int
	PrintValue("Merhaba") // string
	PrintValue(3.14)      // float
}
