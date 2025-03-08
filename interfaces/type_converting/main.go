package main

import "fmt"

func PrintType(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Println("Bu bir string:", v)
	case int:
		fmt.Println("Bu bir integer:", v)
	case bool:
		fmt.Println("Bu bir boolean:", v)
	default:
		fmt.Println("Bilinmeyen tip")
	}
}

func main() {
	PrintType("Hello")  // Bu bir string: Hello
	PrintType(100)      // Bu bir integer: 100
	PrintType(true)     // Bu bir boolean: true
}
