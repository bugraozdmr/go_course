package response

type EpicResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Developer   string `json:"developerDisplayName"`
	Publisher   string `json:"publisherDisplayName"`
	ReleaseDate string `json:"releaseDate"`
	KeyImages   string `json:"Image"`
}