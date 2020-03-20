package artist

import "time"

type UserChat struct {
	ID int `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`

	ArtistID int `gorm:"column:artist_id;not null" json:"artist_id"`
	UserID   int `gorm:"column:user_id;not null" json:"user_id"`

	Comment   string    `gorm:"column:comment;not null" json:"comment"`
	Content   string    `gorm:"column:content" json:"content"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}
