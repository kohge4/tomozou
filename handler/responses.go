package handler

import (
	"time"
)

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

type ChatResponse struct {
	ID        int       `json:"id"`
	Comment   string    `json:"comment"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	UserID    int       `json:"user_id"`
	ArtistID  int       `json:"artist_id"`
	CreatedAt time.Time `json:"created_at"`
}

type MyChatListResponse struct {
	Artists     interface{} `json:"artists"`
	ArtistsInfo interface{} `json:"artists_info"`
}
