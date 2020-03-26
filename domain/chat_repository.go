package domain

type ChatRepository interface {
	SaveChat(*UserChat) error
	ReadChatByArtistID(artistID int) (interface{}, error)
	ReadChatByUserID(userID int) (interface{}, error)
}
