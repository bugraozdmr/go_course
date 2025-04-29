package helpers

import (
	epicFetch "fiyathunter/response/epic"
	gogFetch "fiyathunter/response/gog"

	response "fiyathunter/response"
)

func EpicBindData(data epicFetch.EpicFetch) response.EpicResponse {
	var res response.EpicResponse

	if len(data.Data.Catalog.SearchStore.Elements) == 0 {
		return res
	}

	game := data.Data.Catalog.SearchStore.Elements[0]

	res.Title = game.Title
	res.Description = game.Description
	res.Developer = game.Developer
	res.Publisher = game.Publisher
	res.ReleaseDate = game.ReleaseDate
	res.Slug = game.ProductSlug
	res.Price = game.Price.TotalPrice.FmtPrice.DiscountPrice

	if len(game.KeyImages) > 0 {
		res.KeyImages = game.KeyImages[0].URL
	}

	return res
}

func GogBindData(data gogFetch.GogFetch) response.GogResponse {
	var res response.GogResponse

	if len(data.Products) == 0 {
		return res
	}

	game := data.Products[0]

	res.Title = game.Title
	res.Developers = game.Developers
	res.Publishers = game.Publishers
	res.ReleaseDate = game.ReleaseDate
	res.Price = game.Price.FinalMoney.Amount
	res.Currency = game.Price.FinalMoney.Currency
	res.Discount = game.Price.FinalMoney.Discount
	res.StoreLink = game.StoreLink
	res.Image = game.Image

	return res
}
