package demoapp

//var authenticator rakuten.Authenticator

/*
func init() {
	//authenticator = rakuten.NewAuthenticator("1004273671613614566", "74893f4d646af7296d6d5472d8405e64422c528c", "http://localhost:8000/oauth", "rakuten_favoritebookmark_read")
}

func Login(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, authenticator.AuthURL("state"))
}

func OAuth(c *gin.Context) {
	state := c.Request.URL.Query().Get("state")
	token, err := authenticator.Token(state, c.Request)
	fmt.Printf("okokokok %+v", token)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(token)
	c.SetCookie("rws-token", token.AccessToken, 1000*60*60*24*7, "/", "http://localhost:8000", false, false)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func Index(c *gin.Context) {
	cookie, _ := c.Cookie("rws-token")
	if cookie == "" {
		c.HTML(http.StatusOK, "first.html", nil)
		return
	}
	fmt.Println(cookie)

	c.HTML(http.StatusOK, "index.html", nil)
}

*/
