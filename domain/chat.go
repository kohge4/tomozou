package domain

import "time"

// グループ機能とか作りやすくするために、ToUserID とかは 配列にする
type UserChatIn struct {
	ID      int
	Comment string

	UserID int

	ArtistID []int
	ToUserID []int

	Content   string
	CreatedAt time.Time
}

type UserChat struct {
	ID        int       `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`
	Comment   string    `gorm:"column:comment;not null" json:"comment"`
	UserID    int       `gorm:"column:user_id;not null" json:"user_id"`
	ArtistID  int       `gorm:"column:artist_id;not null" json:"artist_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

/*
func NewChat(userID int, artistID int, comment string) *UserChat {
	return &UserChat{
		ID:       userID,
		ArtistID: artistID,
		Comment:  comment,
	}
}
*/
