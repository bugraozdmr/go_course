package sites

import (
	"fiyathunter/response"
	epic "fiyathunter/sites/epic"
	gog "fiyathunter/sites/gog"
	"fmt"
)

func BestPrice(keyword string) response.BestPrice {
	var response response.BestPrice

	// epic data
	resultEpic := epic.FetchData(keyword)
	response.EpicGames.FinalPrice = fmt.Sprintf("%.2f", float64(resultEpic.PriceDecimal)/100)
	// TODO update discount amount
	response.EpicGames.OriginalPrice = fmt.Sprintf("%.2f", float64(resultEpic.PriceDecimal)/100)
	response.EpicGames.Title = resultEpic.Title
	response.EpicGames.Currency = "TRY"

	// gog data
	resultGog := gog.FetchData(keyword)
	response.GogGames.FinalPrice = resultGog.Price
	// TODO update discount amount
	response.GogGames.OriginalPrice = resultGog.Price
	response.GogGames.Title = resultGog.Title
	response.GogGames.Currency = resultGog.Currency

	return response
}
