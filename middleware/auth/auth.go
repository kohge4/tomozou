package auth

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"tomozou/settings"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const identityKey = "userid"

type User struct {
	UserName    string
	AccountType string
	UserID      int
}

func AuthUser() *jwt.GinJWTMiddleware {
	// spotify 関連の処理を 外部注入できるように
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secretkey"),
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
			println("kekekekekeyyyy")
			println(identityKey)
			println(claims[identityKey])
			return claims[identityKey]
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			userID, _ := c.Get("userid")
			id, ok := userID.(int)

			userName, _ := c.Get("user_name")
			name, ok := userName.(string)
			if ok == false {
				return nil, nil
			}
			return &User{name, "spotify", id}, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			//　得たデータを　実際に 認証する場所
			// MiddlerwareFunc で 認証済みユーザーか否かに使用
			/*
				if v, ok := data.(*User); ok && v.UserName == "admin" {
					return true
				}*/
			//userID, _ := c.Get(identityKey)
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
			id, _ := c.Get("userid")
			tomozouID, _ := c.Get("userid")

			c.JSON(http.StatusOK, gin.H{
				"code":       http.StatusOK,
				"token":      token,
				"expire":     expire.Format(time.RFC3339),
				"id":         id,
				"tomozou_id": tomozouID,
				"url":        settings.FrontURL + "me",
				//"url":        "http://localhost:8080/me",
			})
			// me の エンドポイントを 返せばOK
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
