package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/dghubble/oauth1"
	twauth "github.com/dghubble/oauth1/twitter"

	"tomozou/handler"
	"tomozou/infra/authenticator"
)

const (
	//AuthURL  = "https://api.twitter.com/oauth/authorize"
	AuthURL  = "https://api.twitter.com/oauth/request_token"
	TokenURL = "https://api.twitter.com/oauth2/token"
)
const outOfBand = "oob"

var config oauth1.Config

func main() {
	// oauth2 configures a client that uses app credentials to keep a fresh token

	config2 := &clientcredentials.Config{
		ClientID:     "2U9Fsq4hNOkL9S6dy7awyKrSk",
		ClientSecret: "KbQ6tKJM39HXtD2UptXnAvFpWewFW7WVpTcCP0MJVPkhIJZHTN",
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	config = oauth1.Config{
		//TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
		//ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
		//TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
		ConsumerKey:    "2U9Fsq4hNOkL9S6dy7awyKrSk",
		ConsumerSecret: "KbQ6tKJM39HXtD2UptXnAvFpWewFW7WVpTcCP0MJVPkhIJZHTN",
		CallbackURL:    "https://api.twitter.com/oauth/request_token",
		Endpoint:       twauth.AuthorizeEndpoint,
	}

	/*requestToken, _, err := config.RequestToken()
	if err != nil {
		fmt.Println(err)
	}

	//authorizationURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		fmt.Println(err)
	}
	//http.Redirect(w, req, authorizationURL.String(), http.StatusFound)
	*/

	//func SignIn () {}

	// http.Client will automatically authorize Requests
	httpClient := config2.Client(oauth2.NoContext)

	// Twitter client
	client := twitter.NewClient(httpClient)

	sOpt := &twitter.SearchTweetParams{
		Query: "suchmos",
		Count: 10,
	}
	result, _, _ := client.Search.Tweets(sOpt)

	spotifyHandler := handler.NewSpotifyHandler()

	twitterHandler := handler.NewTwitterHandler()

	fmt.Println(result)

	r := gin.Default()

	//r.LoadHTMLGlob("presentation/views/*")

	crs := cors.DefaultConfig()
	crs.AllowOrigins = []string{"http://localhost:8080", "https://tomozoufront.firebaseapp.com"}
	r.Use(cors.New(crs))

	store := sessions.NewCookieStore([]byte("okokokokoko"))
	// セッションの設定
	// session が cookie を どう触るかわからん
	r.Use(sessions.Sessions("oauth-tweet-test-session", store))

	r.GET("/twi", func(c *gin.Context) {
		c.JSON(200, result)
	})
	r.GET("/twcallback", func(c *gin.Context) {
		// callback 用の URL( "承認しますか" の 後の URL)
		c.JSON(200, result)
	})

	r.GET("/spotify/callback", spotifyHandler.SignUp)
	r.GET("/spotify/login", spotifyHandler.Login)
	r.GET("/spotify/top", func(c *gin.Context) {
		c.String(200, "OkSpotify")
	})
	r.GET("/spotify/me", spotifyHandler.Me)
	r.GET("/spotify/currme", spotifyHandler.MeData)
	r.GET("/spotify/metrack", spotifyHandler.MeTrack)
	r.GET("/spotify/token", spotifyHandler.GetAccessToken)
	r.GET("/spotify/meplaylist", spotifyHandler.MePlaylists)
	r.GET("/spotify/meartists", spotifyHandler.MeArtists)
	r.GET("/spotify/menowplaying", spotifyHandler.MeNowPlaying)
	r.GET("/spotify/merecentlyplaying", spotifyHandler.MeRecentlyPlaying)
	r.GET("/spotify/currentuseralubums", spotifyHandler.CurrentUsersAlbums)

	r.GET("/twitter/callback", twitterHandler.Callback)
	r.POST("/twitter/login", twitterHandler.Login)
	r.GET("/twitter/me", twitterHandler.Me)

	userCtrl := handler.NewUserController()
	r.GET("/profile/user", userCtrl.GetUser)

	// Login 関連 JWT

	authMiddleware := authenticator.Auth()

	r.GET("/cl", handler.ST)
	r.POST("/login", authMiddleware.LoginHandler)
	// ここでの response json は
	// 最終的に LoginResponse=func(c *ginContext) を 実行 , c.JSON(token) 的なやつ

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		// MiddlewareFunc 以下の 関数は Authorization header がないと アクセスできないようになっている
		// MiddlewareFunc の解析順: Jwt から claim を　取ってくる
		// c.Next 以下なので c.Get("JWT_Payload") をかで 値を取ってくる方針
		auth.GET("/hello", func(c *gin.Context) {
			data, _ := c.Get("id")
			c.JSON(200, data)
		})
	}

	r.Run(":8000")
}
