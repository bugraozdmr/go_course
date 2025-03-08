package main

import "fmt"

func main() {
	websites := map[string]string{
		"Youtube" : "https://youtube.com",
		"Amazon" : "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Amazon"])

	websites["Google"] = "https://google.com"
	fmt.Println(websites)

	delete(websites, "Amazon")
	fmt.Println(websites)
}