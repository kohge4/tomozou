package usecase

import "tomozou/domain"

// JWTによるLogin から UserID 取得後にしか 使用できない
type ChatApplication struct {
	ItemRepository domain.ItemRepository
	ChatRepository domain.ChatRepository
}

func (u *ChatApplication) ChatRooms(userID int) (interface{}, error) {
	artists, err := u.ItemRepository.ReadUserArtistTagByUserID(userID)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func (u *ChatApplication) UserComment(chat *domain.UserChat) {
	u.ChatRepository.SaveChat(chat)
	// display chat  が必要
}
