package helpers

import (
	epicFetch "fiyathunter/response/epic"

	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetOfferImageWideURL(images []epicFetch.Image) string {
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
