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

func (repo *ChatDBRepository) ReadChatByUserID(userID int) ([]domain.UserChat, error) {
	return nil, nil
}

func (repo *ChatDBRepository) ReadChatByArtistID(artistID int) ([]domain.UserChat, error) {
	chatList := []domain.UserChat{}
	repo.DB.Where("artist_id = ?", artistID).Find(&chatList)
	return chatList, nil
}
