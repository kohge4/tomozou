package datastore

import "time"

type Artist struct {
	ID int

	SocialID  string
	Image     string
	URL       string
	CreatedAt time.Time
}

type UserChat struct {
	ID int

	ArtistID int
	UserID   int

	Comment   string
	Content   string
	CreatedAt time.Time
}
