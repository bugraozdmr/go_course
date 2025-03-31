package main

import (
	"app/filemanager"
	"app/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// bu dosyadan cekip hesaplar
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// bu ise konsoldan alıyor
		// cmdm := cmdmanager.New()
		// interface kullanarak istedigimiz degerleri alabildik
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("could not process job")
			fmt.Println(err)
		}
	}
}
