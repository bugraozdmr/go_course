package main

import "fmt"

// Interface tanımlıyoruz
type Speaker interface {
	Speak() string
}

// Struct'larımız
type Dog struct{}
type Cat struct{}

// Dog, Speak metodunu uygular
func (d Dog) Speak() string {
	return "Hav hav!"
}

// Cat, Speak metodunu uygular
func (c Cat) Speak() string {
	return "Miyav!"
}

// Interface'i kullanarak fonksiyon yazalım
func MakeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	dog := Dog{}
	cat := Cat{}

	MakeSound(dog) // Hav hav!
	MakeSound(cat) // Miyav!
}
