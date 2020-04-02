package chatdata

import (
	"time"
	"tomozou/domain"
)

// InputされるjsonをBind する 構造体
type ChatIn struct {
	ID      int    `gorm:"column:id;not null;AUTO_INCREMENT" json:"id"`
	Comment string `gorm:"column:comment;not null" json:"comment"`

	UserID int `gorm:"column:user_id;not null" json:"user_id"`

	ArtistID int `gorm:"column:artist_id;not null" json:"artist_id"`
	ToUserID int `gorm:"column:to_user_id" json:"to_user_id"`

	Content   string    `gorm:"column:content" json:"content"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (ch ChatIn) UserChat() (*domain.UserChat, error) {
	chat := &domain.UserChat{
		ID:        ch.ID,
		Comment:   ch.Comment,
		ArtistID:  ch.ArtistID,
		CreatedAt: ch.CreatedAt,
	}
	return chat, nil
}
