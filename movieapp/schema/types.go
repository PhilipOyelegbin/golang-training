package schema

type PosterSizes struct {
	W240 string `json:"w240"`
	W600 string `json:"w600"`
}

type ImageSet struct {
	VerticalPoster PosterSizes `json:"verticalPoster"`
}

type Show struct {
	Id	string `json:"id"`
	Title  string `json:"originalTitle"`
	ReleaseYear   int    `json:"releaseYear"`
	ImageSet ImageSet `json:"imageSet"`
	Overview string `json:"overview"`
	ShowType string `json:"showType"`
	Rating int `json:"rating"`
}

type APIResponse struct {
	Shows []Show `json:"shows"`
}

type HomePageData struct {
	PageTitle string
	Shows     []Show
	CurrentCountry string
}
