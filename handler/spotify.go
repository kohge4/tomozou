package handler

import (
	"fmt"
	"net/http"

	"tomozou/infra/datastore"

	"github.com/gin-gonic/gin"
	"github.com/kohge4/spotify"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	"golang.org/x/oauth2"
)

const (
	//redirectSpotifyURL = "http://localhost:8000/spotify/callback"
	redirectSpotifyURL = "http://localhost:8080/spotify/callback"
	state              = "secret"
	clientID           = "08ad2a3fa89349eabb5b2e9929946b27"
	secretKey          = "10c3f63b95dc4af887ed5f0779a8df6a"
)

type SpotifyHandler struct {
	ClientID    string
	SecretKey   string
	RedirectURL string
	State       string

	Authenticator spotify.Authenticator
	Client        spotify.Client
	Repository    *gorm.DB
	// UserDBRepo を　追加したほうがいい
}

func NewSpotifyHandler() *SpotifyHandler {
	Authenticator := spotify.NewAuthenticator(redirectSpotifyURL, spotify.ScopeUserReadPrivate, spotify.ScopeUserTopRead,
		spotify.ScopeUserReadRecentlyPlayed, spotify.ScopePlaylistModifyPrivate, spotify.ScopePlaylistReadPrivate,
		spotify.ScopePlaylistReadCollaborative, spotify.ScopeUserReadRecentlyPlayed, spotify.ScopeUserReadCurrentlyPlaying)
	return &SpotifyHandler{
		ClientID:    clientID,
		SecretKey:   secretKey,
		RedirectURL: redirectSpotifyURL,
		State:       state,

		Authenticator: Authenticator,
		//Client:        spotify.Client{},
	}
}

func (h *SpotifyHandler) Login(c *gin.Context) {
	//spotifyAuth = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate)
	h.Authenticator.SetAuthInfo(h.ClientID, h.SecretKey)
	//c.Redirect(http.StatusTemporaryRedirect, h.Authenticator.AuthURL(h.State))
	c.JSON(200, Response{200, h.Authenticator.AuthURL(h.State)})
}

func (h *SpotifyHandler) LoginBackend(c *gin.Context) {
	//spotifyAuth = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate)
	h.Authenticator.SetAuthInfo(h.ClientID, h.SecretKey)
	//c.Redirect(http.StatusTemporaryRedirect, h.Authenticator.AuthURL(h.State))
	c.Redirect(200, h.Authenticator.AuthURL(h.State))
}

func (h *SpotifyHandler) Callback(c *gin.Context) {
	// front からの url の prams を使って requets を使うのがベスト？
	token, err := h.Authenticator.Token(h.State, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	//fmt.Println(token)
	//c.SetCookie("spotify-token", token.AccessRoken, )
	h.Client = h.Authenticator.NewClient(token)
	//c.Redirect(http.StatusTemporaryRedirect, "/spotify/me")
	Me, err := h.Client.CurrentUser()
	if err != nil {
		c.String(200, err.Error())
	}
	c.JSON(200, Me)
	fmt.Println("okok")
	//DB に なければ保存する処理？, 毎回 p路フィールで呼ぶのめんどいからよくないかも
	// 最初に ログインしてるなら 云々で処理した方がいいかも
	// cookie に保存する感じな気がする accesstoken を
	userRepo := datastore.NewUserSpotifyDBRepository()
	userS := datastore.UserSpotify{
		UserID:       "dou",
		SocialID:     Me.ID,
		UserName:     Me.DisplayName,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	//userToken := datastore.UserToken{}
	userRepo.Save(userS)
	fmt.Println("Spotify Login Result and Token")
	fmt.Println(userS)
	fmt.Println(token.AccessToken)
	//userSRepo := datastore.NewUserSpotifyDBRepository()

	c.JSON(200, Response{200, "http://localhost:8080/spotify/me"})
}

func (h *SpotifyHandler) Me(c *gin.Context) {
	//fmt.Printf("%T \n", h.Client)
	//認証が住んでないとページに飛べない ==> 何かしらの 前処理ほしい（error 的なの呼べるようにしたい）
	// endpoint /me
	Me, err := h.Client.CurrentUser()
	if err != nil {
		c.String(200, err.Error())
	}
	c.JSON(200, Me)
	fmt.Println("okok")
}

func (h *SpotifyHandler) MeData(c *gin.Context) {
	result, err := h.Client.CurrentUser()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, result)
}

func (h *SpotifyHandler) MeTrack(c *gin.Context) {
	result, err := h.Client.CurrentUsersTracks()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, result)
}

func (h *SpotifyHandler) MeReconmend(c *gin.Context) {
	result, err := h.Client.CurrentUsersTracks()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, result)
}

func (h *SpotifyHandler) MePlaylists(c *gin.Context) {
	result, err := h.Client.GetCurrentUserPlaylist(nil)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, result)
}

// scope の 問題で失敗してる感ある
func (h *SpotifyHandler) MeArtists(c *gin.Context) {
	result, err := h.Client.CurrentUsersTopArtistsOpt(nil)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, result)
}

func (h *SpotifyHandler) CurrentUsersAlbums(c *gin.Context) {
	result, err := h.Client.CurrentUsersAlbumsOpt(nil)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, result)
}

func (h *SpotifyHandler) GetUsersTopArtist(c *gin.Context) {
	result, err := h.Client.CurrentUsersAlbumsOpt(nil)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, result)
}

// DB から accesstoken を拾ってきて, それを用いて spotify api へアクセス
func (h *SpotifyHandler) GetAccessToken(c *gin.Context) {
	userRepo := datastore.NewUserSpotifyDBRepository()
	users := userRepo.ReadAll()
	accessToken := users[0].AccessToken
	fmt.Println(accessToken)
	fmt.Println("refreshTokne")
	fmt.Println(users[0].RefreshToken)
	token := &oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: users[0].RefreshToken,
	}
	h.Client = h.Authenticator.NewClient(token)
	result, err := h.Client.CurrentUser()
	if err != nil {
		fmt.Println(err)
		//h.LoginBackend(c)
		h.Login(c)
		// ここから　再連携しますかにつなげる
	}
	fmt.Println(err)
	c.JSON(200, result)
}
