package httpserver

import (
	"go-auth/internal/app"
	"go-auth/internal/middleware"
	"go-auth/internal/user"
	"net/http"

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
	r.POST("/login", userHandler.Login)

	api := r.Group("/api")

	api.Use(middleware.AuthRequired(a.Config.JWTSecret))

	api.GET("/files", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "This route is protected",
			"files":   []any{},
		})
	})

	api.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success":  true,
			"message":  "This route is protected",
			"products": []any{},
		})
	})

	admin := api.Group("/admin")
	admin.Use(middleware.RequireAdmin())

	admin.GET("/restricted", func(c *gin.Context) {
		role, _ := middleware.GetRole(c)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "This route is restricted",
			"role":    role,
		})
	})
	return r
}
