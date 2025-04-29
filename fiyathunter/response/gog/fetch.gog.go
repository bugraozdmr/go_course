package response

type FinalMoney struct {
	Amount   string `json:"amount"`
	Discount string `json:"discount"`
	Currency string `json:"currency"`
}

type Price struct {
	FinalMoney FinalMoney `json:"finalMoney"`
}

type Products struct {
	Title       string   `json:"title"`
	StoreLink   string   `json:"storeLink"`
	Slug        string   `json:"slug"`
	Price       Price    `json:"price"`
	ReleaseDate string   `json:"releaseDate"`
	Publishers  []string `json:"publishers"`
	Developers  []string `json:"developers"`
	Image       string   `json:"coverHorizontal"`
}

type GogFetch struct {
	Products []Products `json:"products"`
}
