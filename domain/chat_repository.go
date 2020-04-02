package domain

type ChatRepository interface {
	SaveChat(*UserChat) error
	ReadChatByArtistID(artistID int) ([]UserChat, error)
	ReadChatByUserID(userID int) ([]UserChat, error)
}
