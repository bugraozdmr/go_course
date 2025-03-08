package main

import "fmt"

// Boş interface her türü alabilir
func PrintValue(value interface{}) {
	fmt.Println("Değer:", value)
}

func main() {
	PrintValue(42)            // int
	PrintValue("Merhaba!")    // string
	PrintValue(3.14)         // float
	PrintValue(true)         // bool
}
