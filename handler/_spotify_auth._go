package handler

import (
	"fmt"
	"net/http"
	"tomozou/domain"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/oauth2"
)

// User demo
type User struct {
	UserName    string
	AccountType string
	UserID      int
}

const identityKey = "userid"

var userAuthenticator = AuthUser()

func (h *SpotifyHandler) Login(c *gin.Context) {
	//spotifyAuth = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate)
	h.Authenticator.SetAuthInfo(h.ClientID, h.SecretKey)
	//c.Redirect(http.StatusTemporaryRedirect, h.Authenticator.AuthURL(h.State))
	c.JSON(200, Response{200, h.Authenticator.AuthURL(h.State)})
}

// Callback の 代わり
func (h *SpotifyHandler) SignUp(c *gin.Context) {
	token, err := h.Authenticator.Token(h.State, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	h.Client = h.Authenticator.NewClient(token)
	Me, err := h.Client.CurrentUser()
	if err != nil {
		c.String(200, err.Error())
	}

	tomozouUser := domain.User{
		SocialID: Me.ID,
		Name:     Me.DisplayName,
		Auth:     "spotify",
		Image:    Me.Images[0].URL,
	}
	h.UserRepository.Save(tomozouUser)
	uS, _ := h.UserRepository.ReadBySocialID(Me.ID)
	c.Set("tomozou-id", uS[len(uS)-1].ID)
	/*
			fmt.Println("UserData")
			fmt.Println(uS)
			fmt.Println(uS[len(uS)-1].ID)
		test, _ := c.Get("tomozou-id")
		fmt.Println(test)
	*/
	userToken := newUserToken(99, "spotify", token)
	saveUserToken(h.DB, userToken)
	uT := &domain.UserToken{}
	readUserToken(h.DB, uT)

	userAuthenticator.LoginHandler(c)
}

func (h SpotifyHandler) SignIn(c *gin.Context) {
	// ログインや 再連携について　書く？
	fmt.Println(userAuthenticator)
}

func saveUserToken(db *gorm.DB, token *domain.UserToken) {
	if !db.HasTable(&domain.UserToken{}) {
		db.CreateTable(&domain.UserToken{})
	}
	db.Create(token)
}

func readUserToken(db *gorm.DB, token *domain.UserToken) {
	db.First(token)
	fmt.Println(token)
}

func newUserToken(userID int, authType string, token *oauth2.Token) *domain.UserToken {
	return &domain.UserToken{
		UserID:       userID,
		AuthType:     authType,
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}
}
