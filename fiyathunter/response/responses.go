package response

type EpicResponse struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Price        string `json:"price"`
	PriceDecimal int    `json:"-"`
	Developer    string `json:"developerDisplayName"`
	Publisher    string `json:"publisherDisplayName"`
	Slug         string `json:"productSlug"`
	ReleaseDate  string `json:"releaseDate"`
	KeyImages    string `json:"Image"`
}

type GogResponse struct {
	Title       string   `json:"title"`
	StoreLink   string   `json:"storeLink"`
	Price       string   `json:"price"`
	Currency    string   `json:"currency"`
	Discount    string   `json:"discount"`
	Publishers  []string `json:"publishers"`
	Developers  []string `json:"developers"`
	ReleaseDate string   `json:"releaseDate"`
	Image       string   `json:"Image"`
}
