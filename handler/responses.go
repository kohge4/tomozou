package handler

import (
	"time"
	"tomozou/domain"
)

type Response struct {
	Status int    `json:"status"`
	URL    string `json:"url"`
}

type MyProfileResponse struct {
	Me      interface{} `json:"me"`
	Artists interface{} `json:"artists"`
	Tracks  interface{} `json:"tracks"`
}

type MyTrackResponse struct {
}

func NewMyTrackResponse(track *domain.UserTrackTag) *MyTrackResponse {
	return &MyTrackResponse{}
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
