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

	r.GET("/spotify/callback", spotifyHandler.Callback)
	r.GET("/spotify/login", spotifyHandler.Login)
	r.GET("/spotify/top", func(c *gin.Context) {
		c.String(200, "OkSpotify")
	})
	r.GET("/spotify/me", spotifyHandler.Me)

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

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", func(c *gin.Context) {
			c.JSON(200, "okok")
		})
	}

	r.Run(":8000")
}

func login() (requestToken string, err error) {
	requestToken, _, err = config.RequestToken()
	if err != nil {
		return "", err
	}
	authorizationURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		return "", err
	}
	fmt.Printf("Open this URL in your browser:\n%s\n", authorizationURL.String())
	return requestToken, err
}

func receivePIN(requestToken string) (*oauth1.Token, error) {
	fmt.Printf("Paste your PIN here: ")
	var verifier string
	_, err := fmt.Scanf("%s", &verifier)
	if err != nil {
		return nil, err
	}
	// Twitter ignores the oauth_signature on the access token request. The user
	// to which the request (temporary) token corresponds is already known on the
	// server. The request for a request token earlier was validated signed by
	// the consumer. Consumer applications can avoid keeping request token state
	// between authorization granting and callback handling.
	accessToken, accessSecret, err := config.AccessToken(requestToken, "secret does not matter", verifier)
	if err != nil {
		return nil, err
	}
	return oauth1.NewToken(accessToken, accessSecret), err
}
