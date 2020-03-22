package domain

// webservoceaccountImpl が 構造体依存する
type ItemRepository interface {
	SaveUserItem(UserItem) error
	ReadItemByUser(userID int) (interface{}, error)
	ReadArtistBySocialID(socialID string) (*Artist, error)

	SaveArtist(Artist) (int, error)
	SaveTrack(Track) (int, error)
	SaveUserArtistTag(UserArtistTag) error
}
