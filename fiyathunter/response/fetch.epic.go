package response

type FmtPrice struct {
	OriginalPrice string `json:"originalPrice"`
	DiscountPrice string `json:"discountPrice"`
}

type TotalPrice struct {
	OriginalPrice int      `json:"originalPrice"`
	DiscountPrice int      `json:"discountPrice"`
	FmtPrice      FmtPrice `json:"fmtPrice"`
}

type Price struct {
	TotalPrice TotalPrice `json:"totalPrice"`
}

type Image struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Game struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       Price  `json:"price"`
	Developer   string `json:"developerDisplayName"`
	Publisher   string `json:"publisherDisplayName"`
	ReleaseDate string `json:"releaseDate"`
	KeyImages   []Image `json:"keyImages"`
}

type SearchStore struct {
	Elements []Game `json:"elements"`
}

type Catalog struct {
	SearchStore SearchStore `json:"searchStore"`
}

type Data struct {
	Catalog Catalog `json:"Catalog"`
}

type EpicFetch struct {
	Data Data `json:"data"`
}
