package handler

import "github.com/gin-gonic/gin"

func (u *UserProfileApplicationImpl) Debug(c *gin.Context) {
	a := u.Handler.DebugItem(1)
	c.JSON(200, a)
}
