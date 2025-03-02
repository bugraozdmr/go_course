package utils

import (
	"fmt"
	"log"
	"os"
)

// Bakiye dosyasını okur, eğer boşsa 1000 ile başlatır.
func ReadBalanceFromFile(filename string) float64 {
	var accountBalance float64

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = fmt.Fscanf(file, "%f", &accountBalance)
	if err != nil {
		accountBalance = 1000
	}
	return accountBalance
}

// Bakiye dosyasını günceller.
func writeBalanceToFile(filename string, balance float64) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%f", balance))
	if err != nil {
		log.Fatal(err)
	}
}
