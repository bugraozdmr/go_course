package main

import "fmt"

// `T` adlı bir generic parametre tanımladık
// T any her türü alabilen bir generic tür anlamına gelir.
// Print fonksiyonu farklı tipler için tekrar yazılmak zorunda değil!
func Print[T any](value T) {
	fmt.Println(value)
}

func main() {
	Print(42)         // int
	Print("Hello")    // string
	Print(3.14)       // float
	Print(true)       // bool
}
