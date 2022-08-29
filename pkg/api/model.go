package api

type CandyStore struct {
	Name  string
	Candy string
	Eaten int
}

type FavoriteSnack struct {
	Name          string `json:"name"`
	FavoriteSnack string `json:"favoriteSnack"`
	TotalSnacks   int    `json:"totalSnacks"`
}
