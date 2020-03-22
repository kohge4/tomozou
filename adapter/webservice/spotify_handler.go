package webservice

import (
	"tomozou/domain"

	"github.com/jinzhu/gorm"
	"github.com/kohge4/spotify"

	_ "github.com/mattn/go-sqlite3"
)

const (
	//redirectSpotifyURL = "http://localhost:8000/spotify/callback"
	redirectSpotifyURL = "http://localhost:8080/spotify/callback"
	state              = "secret"
	clientID           = "08ad2a3fa89349eabb5b2e9929946b27"
	secretKey          = "10c3f63b95dc4af887ed5f0779a8df6a"
)

type SpotifyHandler struct {
	domain.WebServiceAccount
	ClientID    string
	SecretKey   string
	RedirectURL string
	State       string

	Authenticator spotify.Authenticator
	Client        spotify.Client
	DB            *gorm.DB

	UserRepository    domain.UserRepository
	SpotifyRepository domain.ItemRepository
}

//  認証　にも使用したいから domain.WebServiceAccount にしない
func NewSpotifyHandler(userRepo domain.UserRepository, spRepo domain.ItemRepository, db *gorm.DB) *SpotifyHandler {
	Authenticator := spotify.NewAuthenticator(redirectSpotifyURL, spotify.ScopeUserReadPrivate, spotify.ScopeUserTopRead,
		spotify.ScopeUserReadRecentlyPlayed, spotify.ScopePlaylistModifyPrivate, spotify.ScopePlaylistReadPrivate,
		spotify.ScopePlaylistReadCollaborative, spotify.ScopeUserReadRecentlyPlayed, spotify.ScopeUserReadCurrentlyPlaying)

	return &SpotifyHandler{
		ClientID:    clientID,
		SecretKey:   secretKey,
		RedirectURL: redirectSpotifyURL,
		State:       state,

		Authenticator: Authenticator,

		UserRepository:    userRepo,
		SpotifyRepository: spRepo,
		DB:                db,
	}
}

// 認証が 終わった後は こっちにして UseCase の実行をしていく
func (h *SpotifyHandler) ConvertWebServiceAccountImpl() domain.WebServiceAccount {
	return h
}

func (h *SpotifyHandler) User() (*domain.User, error) {
	me, err := h.Client.CurrentUser()
	if err != nil {
		return nil, err
	}
	user := domain.User{
		SocialID: me.ID,
		Name:     me.DisplayName,
		Auth:     "spotify",
		Image:    me.Images[0].URL,
	}
	return &user, nil
}

func (h *SpotifyHandler) SaveUserItem(userID int) error {
	h.saveUserToken(userID)

	h.saveTopArtists(userID)
	h.saveRecentlyFavoriteArtists(userID)
	//h.saveRecentlyPlayedTracks(userID)
	//h.saveTopTracks()
	//h.saveNowPlayingTrack()
	return nil
}
