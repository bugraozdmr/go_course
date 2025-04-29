package analyzers

import (
	"fmt"
	"strconv"
)

type GamePrice struct {
	Title         string
	OriginalPrice string
	FinalPrice    string
	Currency      string
}

func ComparePrices(epic GamePrice, gog GamePrice) string {
	epicPrice, err1 := strconv.ParseFloat(epic.FinalPrice, 64)
	gogPrice, err2 := strconv.ParseFloat(gog.FinalPrice, 64)

	if err1 != nil || err2 != nil {
		return "Hatalı fiyat formatı"
	}

	// USD'yi TRY'ye çevir
	if gog.Currency == "USD" {
		gogPrice *= 40
	}

	result := ""
	if epicPrice < gogPrice {
		result = fmt.Sprintf("Epic Games daha ucuz: %.2f %s < %.2f TRY", epicPrice, epic.Currency, gogPrice)
	} else if epicPrice > gogPrice {
		result = fmt.Sprintf("GOG daha ucuz: %.2f TRY < %.2f %s", gogPrice, epicPrice, epic.Currency)
	} else {
		result = "İki fiyat eşit"
	}

	return result
}
