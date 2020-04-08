package domain

type Track struct {
	ID       int    `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`
	SocialID string `gorm:"column:name;not null" json:"name"`
	Name     string
	TrackURL string

	ArtistName string
	ArtistID   string
}
