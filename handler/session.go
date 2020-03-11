package handler

import "github.com/gin-gonic/contrib/sessions"

var (
	store        = sessions.NewCookieStore([]byte("secret"))
	sessionNames = []string{"twitter-auth", "spotify-auth"}
)
