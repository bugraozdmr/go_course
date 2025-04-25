package helpers

import (
	"fiyathunter/response"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetOfferImageWideURL(images []response.Image) string {
	for _, img := range images {
		if img.Type == "OfferImageWide" {
			return img.URL
		}
	}
	return ""
}

func GetUrl(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	epicURL := os.Getenv("EPIC_FULL_URL")

	return epicURL
}

func EpicBindData(data response.EpicFetch) response.EpicResponse {
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
	res.Price = game.Price.TotalPrice.FmtPrice.DiscountPrice

	if len(game.KeyImages) > 0 {
		res.KeyImages = game.KeyImages[0].URL
	}

	return res
}
