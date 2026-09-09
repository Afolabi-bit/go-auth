package httpserver

import (
	"go-auth/internal/app"
	"go-auth/internal/user"

	"github.com/gin-gonic/gin"
)

func NewRouter(a *app.App) *gin.Engine {
	r := gin.New()
	r.HandleMethodNotAllowed = true
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", health)

	// user routes
	userRepo := user.NewRepo(a.MongoDatabase)
	userSvc := user.NewService(userRepo, a.Config.JWTSecret)
	userHandler := user.NewHandler(userSvc)

	r.POST("/register", userHandler.Register)

	return r
}
