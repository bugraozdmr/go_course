package main

import "fmt"

// Generic fonksiyon: Map içinde key'e göre arama yapar
func FindInMap[K comparable, V any](m map[K]V, key K) (V, bool) {
	value, found := m[key]
	return value, found
}

func main() {
	// String - int map
	ages := map[string]int{"Ali": 30, "Veli": 25}
	age, found := FindInMap(ages, "Ali")
	fmt.Println(age, found) // 30 true

	// Int - string map
	names := map[int]string{1: "Ahmet", 2: "Mehmet"}
	name, found := FindInMap(names, 2)
	fmt.Println(name, found) // Mehmet true
}
