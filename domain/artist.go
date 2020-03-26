package domain

import "time"

type Artist struct {
	ID int `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`

	Name      string    `gorm:"column:name;not null" json:"name"`
	SocialID  string    `gorm:"column:social_id;not null" json:"social_id"`
	Image     string    `gorm:"column:image;not null" json:"image"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func NewArtist(name string, socialID string, image string) Artist {
	return Artist{
		Name:     name,
		SocialID: socialID,
		Image:    image,
	}
}

type ArtistTrackTag struct {
	ID int

	ArtistID int
	TrackID  int
	TagName  string
}
