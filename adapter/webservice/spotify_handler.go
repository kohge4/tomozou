package webservice

import (
	"tomozou/domain"
	"tomozou/settings"

	"github.com/jinzhu/gorm"
	"github.com/kohge4/spotify"

	_ "github.com/mattn/go-sqlite3"
)

const (
	redirectSpotifyURL = settings.FrontURL + "spotify/callback"
	//redirectSpotifyURL = "http://localhost:8080/spotify/callback"
	state     = "secret"
	clientID  = "08ad2a3fa89349eabb5b2e9929946b27"
	secretKey = "10c3f63b95dc4af887ed5f0779a8df6a"
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

	var image string
	images := me.Images
	if len(images) == 0 {
		image = "https://1.bp.blogspot.com/-ytFmslbH_nw/VufYbH7dO9I/AAAAAAAA43w/ds0JuKtPQVcai8to9nUb77g0pIp8iOT_w/s800/music_norinori_woman.png"
	} else {
		image = images[0].URL
	}

	user := domain.User{
		SocialID: me.ID,
		Name:     me.DisplayName,
		Auth:     "spotify",
		Image:    image,
	}
	return &user, nil
}

func (h *SpotifyHandler) SaveUserItem(userID int) error {
	h.saveUserToken(userID)

	h.saveTopArtists(userID)
	h.saveRecentlyFavoriteArtists(userID)
	//h.saveRecentlyPlayedTracks(userID)
	h.saveTopTracks(userID)
	h.saveNowPlayingTrack(userID)
	return nil
}

func (h *SpotifyHandler) UpdateUserItem(userID int) error {
	// 最初に削除の コードを書く
	// 大元の ID を 保持するか迷うね
	h.updateUserToken(userID)

	h.deleteUserArtistInfo(userID)
	h.saveTopArtists(userID)
	h.saveRecentlyFavoriteArtists(userID)

	h.saveTopTracks(userID)
	h.saveNowPlayingTrack(userID)
	return nil
}

// このopt は domain から 型を用いた方がいい可能性あり
func (h *SpotifyHandler) UpdateUserItemOpt(userID int, opt string) error {
	switch opt {
	case "nowplaying":
		h.saveNowPlayingTrack(userID)
	default:
	}
	return nil
}

func (h *SpotifyHandler) DebugItem(userID int) interface{} {
	/*
		//h.saveTopTracks(userID)
		timerange := "short"
		limit := 2
		opt := &spotify.Options{
			Timerange: &timerange,
			Limit:     &limit,
		}
		// options がうまくいってなかった
		results, err := h.Client.GetUserTopTracks2Opt(opt)
		if err != nil {
			return err
		}
		fmt.Println(*opt)
	*/
	/*
		results, err := h.Client.PlayerRecentlyPlayed()
		if err != nil {
			return err
		}
	*/
	var track *spotify.SimpleTrack

	results, err := h.Client.PlayerCurrentlyPlaying()
	if err != nil {
		return err
	}
	if results.Item == nil {
		recentlyPlaying, err := h.Client.PlayerRecentlyPlayed()
		if err != nil {
			return err
		}
		track = &recentlyPlaying[0].Track
	} else {
		nowPlaying, err := h.Client.PlayerCurrentlyPlaying()
		if err != nil {
			return err
		}
		track = nowPlaying.Item.ToSimpleTrack()
	}

	return track
}
