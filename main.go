package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"tomozou/adapter/webservice"
	"tomozou/controller"
	"tomozou/handler"
	"tomozou/infra/datastore"
	"tomozou/usecase"
)

func main() {

	gormConn, _ := datastore.GormConn()
	userRepo := datastore.NewUserDBRepository(gormConn)
	itemRepo := datastore.NewSpotifyItemDBRepository(gormConn)

	useCase := usecase.NewUserProfileApplication(userRepo, itemRepo)

	spotifyHandler := webservice.NewSpotifyHandler(userRepo, itemRepo, gormConn)
	authMiddleware := handler.AuthUser()

	userProfileAppImpl := controller.UserProfileApplicationImpl{
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
	r.GET("/spotify/me", userProfileAppImpl.Me)

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", func(c *gin.Context) {
			data, _ := c.Get("id")
			c.JSON(200, data)
		})
	}
	r.Run(":8000")

}
