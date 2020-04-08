package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getIDFromContext(c *gin.Context) (int, error) {
	id, _ := c.Get("userid")
	userID, ok := id.(float64)
	if ok == false {
		c.String(403, "Authentication is failed")
		return 0, fmt.Errorf("Authentication is failed")
	}
	if userID == 0 {
		c.String(403, "Authentication is failed")
		return 0, fmt.Errorf("Authentication is failed")
	}
	return int(userID), nil
}

func embedTrackURLFromSopotifyID() {}

func embedArtistURLFromSopotifyID() {}

func embedTrackURLFromAppleID() {}

func embedArtistURLFromAppleID() {}
