package main

import "fmt"

// tek bir değer dolaşıyor uygulamada

func main() {
	age := 32 // regular variable

	// baska tanimlama
	// var agePointer *int
	agePointer := &age

	fmt.Println("Age pointer address : ", agePointer)
	fmt.Println("Age value : ", *agePointer)
	fmt.Println("Age: ",age)

	// adultYears := getAdultYears(agePointer)
	// fmt.Println("Adult Years: ", adultYears)

	getAdultYears(agePointer)
	fmt.Println("Adult Years: ", *agePointer)
}

// 2 tane age olusturuyor garbage collector caliscak -- onun yerine pointer istiyoruz
func getAdultYears(age *int) {
	// return *age - 18

	// direkt esleme yapıyoruz return'a gerek yok
	*age = *age - 18
}