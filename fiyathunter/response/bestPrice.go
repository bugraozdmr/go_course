package response

type EpicGames struct {
	Title         string `json:"title"`
	OriginalPrice string `json:"originalPrice"`
	FinalPrice    string `json:"finalPrice"`
	Currency      string `json:"currency"`
}

type GogGames struct {
	Title         string `json:"title"`
	OriginalPrice string `json:"originalPrice"`
	FinalPrice    string `json:"finalPrice"`
	Currency      string `json:"currency"`
}

type BestPrice struct {
	EpicGames EpicGames `json:"epic"`
	GogGames  GogGames  `json:"gog"`
}
