package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"tomozou/infra/datastore"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// User demo
type User struct {
	UserName    string
	AccountType string
	UserID      int
}

const identityKey = "userid"

func AuthUser() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "userid",
		// オリジナルで追加.... scope や permission を　追加したい
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					"userid": v.UserID,
					"id":     v.UserID,
					"name":   v.UserName,
					"scope":  []string{v.AccountType},
					"login":  v.AccountType,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			// jwt の claims から user name を　取ってくる
			claims := jwt.ExtractClaims(c)
			c.Set(identityKey, claims[identityKey])
			//claims[identityKey].(int)
			return claims[identityKey]
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			userID, _ := c.Get("tomozou-id")
			id, ok := userID.(int)
			fmt.Println(id)
			println(ok)
			if ok == false {
				return nil, nil
			}
			return &User{"test", "spotify", id}, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			//　得たデータを　実際に 認証する場所
			// MiddlerwareFunc で 認証済みユーザーか否かに使用
			/*
				if v, ok := data.(*User); ok && v.UserName == "admin" {
					return true
				}*/
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			/* httprequest の token をひっぱてくるコード
			 => よって ここでは nil
			claims := jwt.ExtractClaims(c)
			c.Set(identityKey, claims[identityKey])
			id, _ := c.Get(identityKey)
			*/
			fmt.Println(token)
			id, _ := c.Get("tomozou-id")

			tomozouID, _ := c.Get("tomozou-id")
			c.JSON(http.StatusOK, gin.H{
				"code":       http.StatusOK,
				"token":      token,
				"expire":     expire.Format(time.RFC3339),
				"ID":         id,
				"tomozou-id": tomozouID,
			})
		},

		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}

var userAuthenticator = AuthUser()

func (h SpotifyHandler) SignUp(c *gin.Context) {
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

	userRepo := datastore.NewUserSpotifyDBRepository()
	userS := datastore.UserSpotify{
		UserID:       "dou",
		SocialID:     Me.ID,
		UserName:     Me.DisplayName,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	// UserDB から すでに 使用した ID かをチェックする関数欲しいね

	//userToken := datastore.UserToken{}
	userRepo.Save(userS)
	fmt.Println("Spotify Login Result and Token")
	fmt.Println(userS)
	fmt.Println(token.AccessToken)
	//userSRepo := datastore.NewUserSpotifyDBRepository()

	//c.JSON(200, Response{200, "http://localhost:8080/spotify/me"})

	tomozouUserRepo := datastore.NewUserDBRepository()

	tomozouUser := datastore.User{
		SocialID: Me.ID,
		Name:     Me.DisplayName,
		Auth:     "spotify",
		Image:    Me.Images[0].URL,
	}

	tomozouUserRepo.Save(tomozouUser)
	uS := tomozouUserRepo.ReadBySocialID(Me.ID)
	c.Set("tomozou-id", uS[len(uS)-1].ID)
	fmt.Println("UserData")
	fmt.Println(uS)
	fmt.Println(uS[len(uS)-1].ID)
	test, _ := c.Get("tomozou-id")
	fmt.Println(test)

	userAuthenticator.LoginHandler(c)
}

func (h SpotifyHandler) SignIn(c *gin.Context) {
	fmt.Println(userAuthenticator)
}
