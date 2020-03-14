package handler

import (
	"fmt"
	"tomozou/domain"
	"tomozou/infra/datastore"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserRepository domain.UserRepository
}

func NewUserController() *UserController {
	return &UserController{}
}

func (ctrl *UserController) Me(c *gin.Context) {
	userRepo := datastore.NewUserDBRepository()
	users := userRepo.ReadAll()
	c.JSON(200, users[0])
}

func (ctrl *UserController) GetUser(c *gin.Context) {
	// context の url の　param を用いて使用
	userRepo := datastore.NewUserDBRepository()
	users := userRepo.ReadAll()
	c.JSON(200, users[0])
}

// いか　テスト用

func ST(c *gin.Context) {
	userRepo := datastore.NewUserSpotifyDBRepository()
	users := userRepo.ReadAll()
	fmt.Println("!accesstoken")
	fmt.Println(users[0].AccessToken)
	fmt.Println("refreshtoken")
	fmt.Println(users[0].RefreshToken)
	fmt.Println(users[0])
	c.JSON(200, users[0].AccessToken)
}
