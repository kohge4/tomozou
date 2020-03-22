package handler

import (
	"fmt"
	"tomozou/domain"
	"tomozou/infra/datastore"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var gormConn *gorm.DB

type UserController struct {
	UserRepository domain.UserRepository
	// DB *gorm.DB も いけるかも
	// domain の repository を 扱えない問題
}

func NewUserController(userRepo domain.UserRepository) *UserController {
	return &UserController{
		UserRepository: userRepo,
	}
}

func (ctrl *UserController) Me(c *gin.Context) {
	userRepo := datastore.NewUserDBRepository(gormConn)
	users := userRepo.ReadAll()
	c.JSON(200, users[0])
}

func (ctrl *UserController) GetUser(c *gin.Context) {
	// context の url の　param を用いて使用
	userRepo := datastore.NewUserDBRepository(gormConn)
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
