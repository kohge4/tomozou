package authenticator

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	// 同一 context 内だから middleware で 作った　値がそのまま使える
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*User).UserName,
		"text":     "Hello World.",
	})
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func Auth() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		// オリジナルで追加.... scope や permission を　追加したい
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			scope := []string{"spotify"}
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
					// ここで Spotify とか twitter  とか増やしていける
					// ここの 値 spotify: "", twitter: "とかやってもいいかもね" ==> 容量的にきついかも
					"scope": scope,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			// jwt の claims から user name を　取ってくる
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			// Accesstoken を　取ってくることに成功したら
			// これの 戻り値　が data

			/*方針
			AuthHeader を確認 => mw.TokenLookUp で いける
			parsetoken かも

			*/

			/*
					var loginVals login
					if err := c.ShouldBind(&loginVals); err != nil {
						return "", jwt.ErrMissingLoginValues
					}
					userID := loginVals.Username
					password := loginVals.Password

					if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
						return &User{
							UserName:  userID,
							LastName:  "Bo-Yi",
							FirstName: "Wu",
						}, nil
					}
				return nil, jwt.ErrFailedAuthentication
			*/
			return &User{
				UserName: "testdemo",
			}, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			//　得たデータを　実際に 認証する場所
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		/* 必要追記: IdentityHandler は Userを識別 func(c *gin.Context) の形式
		ExtractClaims をよんでいる
			読み込んだ payload を context."JWT_PAYLOAD" に保存しているので, それを読み込む


		*/

		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		// 複数の 可能性を 書いておくことができる (3 箇所のうちどこかに JWT が存在すれば良い )
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
