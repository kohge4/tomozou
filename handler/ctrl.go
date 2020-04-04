package handler

import (
	"fmt"
	"net/http"
	"strconv"

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
	accessToken, err := u.Handler.Authenticator.Token(u.Handler.State, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	u.Handler.Client = u.Handler.Authenticator.NewClient(accessToken)

	// ここで UseCase に切り替える
	u.UseCase.WebServiceAccount = u.Handler.ConvertWebServiceAccountImpl()

	existingUser, err := u.UseCase.CheckExistingUser()
	if err != nil {
		c.String(403, err.Error())
	}
	if existingUser != nil {
		// すでに そのサービスでログインしたことあるユーザーの場合
		c.Set("userid", existingUser.ID)
		c.Set("user_name", existingUser.Name)

		err = u.UseCase.UpdateUser(existingUser.ID)
		if err != nil {
			c.String(403, err.Error())
		}
		u.AuthMiddleware.LoginHandler(c)
		return
	}

	user, err := u.UseCase.RegistryUser()
	if err != nil {
		c.String(403, err.Error())
	}
	c.Set("userid", user.ID)
	c.Set("user_name", user.Name)
	/*
		fmt.Println("UserCheck")
		println(user.Name)
		println(user.ID)
	*/
	u.AuthMiddleware.LoginHandler(c)
}

func (u *UserProfileApplicationImpl) MyProfile(c *gin.Context) {
	id, _ := c.Get("userid")
	userID, ok := id.(float64)
	if ok == false {
		c.String(403, "Authentication is failed")
	}
	if userID == 0 {
		c.String(403, "Authentication is failed")
	}
	me, err := u.UseCase.Me(int(userID))
	if err != nil {
		c.String(403, err.Error())
	}
	tag, err := u.UseCase.MyUserArtistTag(int(userID))
	if err != nil {
		return
	}

	response := MyProfileResponse{
		Me:      me,
		Artists: tag,
	}
	c.JSON(200, response)
}

func (u *UserProfileApplicationImpl) MyChatList(c *gin.Context) {
	id, _ := c.Get("tomozou-id")
	userID, _ := id.(int)

	tag, err := u.UseCase.MyUserArtistTag(userID)
	if err != nil {
		return
	}
	response := MyChatListResponse{
		Artists:     tag,
		ArtistsInfo: "",
	}
	c.JSON(200, response)
}

func (u *UserProfileApplicationImpl) Me(c *gin.Context) {
	id, _ := c.Get("userid")
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
	println(userID)
	c.JSON(200, myArtists)
}

func (u *UserProfileApplicationImpl) SearchUsersByArtistID(c *gin.Context) {
	artistID := c.Param("artistID")
	id, _ := strconv.Atoi(artistID)
	fmt.Println(id)

	users, err := u.UseCase.DisplayUsersByArtistID(id)
	if err != nil {
		c.JSON(403, err.Error())
	}
	c.JSON(200, users)
}

func (u *UserProfileApplicationImpl) SearchUsersByArtistName(c *gin.Context) {
	name := c.Query("name")
	users, err := u.UseCase.DisplayUsersByArtistName(name)
	if err != nil {
		c.JSON(403, err.Error())
	}
	c.JSON(200, users)
}

func (u *UserProfileApplicationImpl) MyTrack(c *gin.Context) {

}
