package datastore

import (
	"tomozou/domain"

	"github.com/jinzhu/gorm"
)

type ChatDBRepository struct {
	DB *gorm.DB
}

func NewChatDBRepository(db *gorm.DB) domain.ChatRepository {
	return &ChatDBRepository{
		DB: db,
	}
}

func (repo *ChatDBRepository) SaveChat(chat *domain.UserChat) error {
	repo.DB.Create(chat)
	return nil
}

func (repo *ChatDBRepository) ReadChatByUserID(userID int) (interface{}, error) {
	return nil, nil
}

func (repo *ChatDBRepository) ReadChatByArtistID(artistID int) (interface{}, error) {
	return nil, nil
}
