package handler

import (
	"fmt"
	"net/http"

	"tomozou/infra/datastore"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify"
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
	// UserDBRepo を　追加したほうがいい
}

func NewSpotifyHandler() *SpotifyHandler {
	Authenticator := spotify.NewAuthenticator(redirectSpotifyURL, spotify.ScopeUserReadPrivate)
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
	Me, err := h.Client.CurrentUser()
	if err != nil {
		c.String(200, err.Error())
	}
	c.JSON(200, Me)
	fmt.Println("okok")
	//DB に なければ保存する処理？, 毎回 p路フィールで呼ぶのめんどいからよくないかも
	// 最初に ログインしてるなら 云々で処理した方がいいかも
	// cookie に保存する感じな気がする accesstoken を
	/*
		userRepo := datastore.NewUserDBRepository()
		user := datastore.User{
			SocialID: Me.ID,
			Name:     Me.DisplayName,
			Auth:     "ok",
		}
		userRepo.Save(user)
		users := userRepo.ReadAll()
		fmt.Println(users)
	*/
}
