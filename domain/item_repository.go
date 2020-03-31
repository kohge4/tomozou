package domain

// webservoceaccountImpl が 構造体依存する
type ItemRepository interface {
	ReadItemByUser(userID int) (interface{}, error)
	ReadArtistBySocialID(socialID string) (*Artist, error)
	SaveArtist(Artist) (int, error)
	SaveTrack(Track) (int, error)
	SaveUserArtistTag(UserArtistTag) error

	ReadTagByUser(userID int) (interface{}, error)
	ReadUserArtistTagByUserID(userID int) (interface{}, error)

	ReadUserIDByArtistID(artistID int) ([]int, error)
}
