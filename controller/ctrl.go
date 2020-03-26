package controller

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"tomozou/adapter/webservice"
	"tomozou/usecase"
)

type UserProfileApplicationImpl struct {
	UseCase *usecase.UserProfileApplication

	Handler        *webservice.SpotifyHandler
	AuthMiddleware *jwt.GinJWTMiddleware
}

func (u *UserProfileApplicationImpl) Login(c *gin.Context) {
	u.Handler.Authenticator.SetAuthInfo(u.Handler.ClientID, u.Handler.SecretKey)
	c.JSON(200, Response{200, u.Handler.Authenticator.AuthURL(u.Handler.State)})
}

func (u *UserProfileApplicationImpl) Callback(c *gin.Context) {
	// Login が成功したら UserCase の domain.WebSeerviceAccount を更新する
	// => 更新してから RegistryUserを実行する
	token, err := u.Handler.Authenticator.Token(u.Handler.State, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	u.Handler.Client = u.Handler.Authenticator.NewClient(token)

	u.UseCase.WebServiceAccount = u.Handler.ConvertWebServiceAccountImpl()
	u.UseCase.RegistryUser()

	u.AuthMiddleware.LoginHandler(c)
}

func (u *UserProfileApplicationImpl) MyProfile(c *gin.Context) {
	id, _ := c.Get("tomozou-id")
	userID, _ := id.(int)

	if userID == 0 {
		userID = 1
	}
	me, err := u.UseCase.Me(userID)
	if err != nil {
		c.String(403, err.Error())
	}
	tag, err := u.UseCase.MyUserArtistTag(userID)
	if err != nil {
		return
	}

	response := MyProfileResponse{
		Me:      me,
		Artists: tag,
		//TopArtists:      tag,
		//FavoriteArtists: tag,
	}
	c.JSON(200, response)
}

func (u *UserProfileApplicationImpl) Me(c *gin.Context) {
	id, _ := c.Get("tomozou-id")
	userID, _ := id.(int)
	me, err := u.UseCase.DisplayMe(userID)
	if err != nil {
		c.String(403, err.Error())
	}
	c.JSON(200, me)
}

func (u *UserProfileApplicationImpl) MyArtist(c *gin.Context) {
	id, _ := c.Get("tomozou-id")
	userID, _ := id.(int)
	myArtists, err := u.UseCase.MyArtistTag(1)
	if err != nil {
		c.JSON(403, err.Error())
	}
	println("CCHHH")
	println(id)
	println(userID)
	c.JSON(200, myArtists)
}

func (u *UserProfileApplicationImpl) MyTrack(c *gin.Context) {

}
