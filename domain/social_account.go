package domain

type WebServiceAccount interface {
	User() (*User, error)
	Content() (UserItem, error)
	Link(User) error

	// SpotifyHandler.SaveUserItem で 必要な情報を 全部保存
	SaveUserItem(userID int) error

	SignIn() (string, error)
	Callback() (string, error)
	SignUp() (string, error)
	SignOut(User) error
	//DeleteUser(User) error
}

type UserItem struct {
	TopArtist              []Artist
	RecentlyFavoriteArtist []Artist
	TopTrack               []Track
	RecentlyPlayedTrack    []Track
}
