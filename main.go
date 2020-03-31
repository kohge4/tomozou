package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"tomozou/adapter/webservice"
	"tomozou/domain"
	"tomozou/handler"
	"tomozou/infra/datastore"
	"tomozou/middleware/auth"
	"tomozou/usecase"
)

func main() {

	gormConn, _ := datastore.GormConn()
	userRepo := datastore.NewUserDBRepository(gormConn)
	itemRepo := datastore.NewSpotifyItemDBRepository(gormConn)

	useCase := usecase.NewUserProfileApplication(userRepo, itemRepo)

	spotifyHandler := webservice.NewSpotifyHandler(userRepo, itemRepo, gormConn)
	authMiddleware := auth.AuthUser()

	userProfileAppImpl := handler.UserProfileApplicationImpl{
		UseCase: useCase,

		Handler:        spotifyHandler,
		AuthMiddleware: authMiddleware,
	}

	r := gin.Default()

	crs := cors.DefaultConfig()
	crs.AllowOrigins = []string{"http://localhost:8080", "https://tomozoufront.firebaseapp.com"}
	r.Use(cors.New(crs))

	r.GET("/spotify/callback", userProfileAppImpl.Callback)
	r.GET("/spotify/login", userProfileAppImpl.Login)
	r.GET("/spotify/top", func(c *gin.Context) {
		c.String(200, "OkSpotify")
	})
	r.GET("/spotify/me", userProfileAppImpl.MyProfile)
	r.GET("/spotify/myartist", userProfileAppImpl.MyArtist)
	r.GET("/search/user/byartist", userProfileAppImpl.SearchUsersByArtist)

	auth := r.Group("/me")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", func(c *gin.Context) {
			data, _ := c.Get("id")
			c.JSON(200, data)
		})
		auth.GET("/profile", userProfileAppImpl.MyProfile)
	}

	devUserRepo := datastore.NewDevUserRepo(gormConn)
	r.GET("/dev/user", func(c *gin.Context) {
		users, _ := devUserRepo.CheckUser()
		c.JSON(200, users)
	})
	r.GET("/dev/tag", func(c *gin.Context) {
		tags := []domain.UserArtistTag{}
		devUserRepo.DB.Find(&tags)
		c.JSON(200, tags)
	})
	r.GET("/dev/userdata", func(c *gin.Context) {
	})

	// Chat 用: authによるJWT 以下から
	chatRepo := datastore.NewChatDBRepository(gormConn)
	chatApp := usecase.ChatApplication{
		ItemRepository: itemRepo,
		ChatRepository: chatRepo,
	}
	chatAppImpl := handler.ChatApplicationImpl{
		UseCase: chatApp,
	}
	r.GET("/chat/room", chatAppImpl.DisplayChatRoom)
	r.POST("/chat/user/comment", chatAppImpl.UserChat)

	r.Run(":8000")
}
