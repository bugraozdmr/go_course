package main

import "fmt"

// Generic struct: `T` farklı türler olabilir
type Box[T any] struct {
	Value T
}

func main() {
	intBox := Box[int]{Value: 100}
	stringBox := Box[string]{Value: "Merhaba!"}

	fmt.Println(intBox.Value)  // 100
	fmt.Println(stringBox.Value) // Merhaba!
}
