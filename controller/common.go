package controller

type Response struct {
	Status int    `json:"status"`
	URL    string `json:"url"`
}

type MyProfileResponse struct {
	Me      interface{} `json:"me"`
	Artists interface{} `json:"artists"`
	//TopArtists      interface{} `json:"top_artists"`
	//FavoriteArtists interface{} `json:"favorite_artists"`
}
