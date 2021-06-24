package main

import (
	"assessment/auth"
	"assessment/config"
	"assessment/handler"
	"assessment/site"
	"assessment/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB = config.Config()

	userRepository = user.NewRepo(DB)
	userService    = user.NewService(userRepository, authService)
	authService    = auth.NewAuthService()
	userHandler    = handler.NewUserHandler(userService, authService)
	middleware     = handler.Middleware(userService, authService)

	siteRepository = site.NewRepo(DB)
	siteService    = site.NewService(siteRepository)
	siteHandler    = handler.NewSiteHandler(siteService)
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	// r.GET("/users", middleware, userHandler.GetAllUser)
	r.POST("/users/register", userHandler.RegisterUser)
	r.POST("/users/login", userHandler.LoginUser)
	r.GET("users/:user_id", middleware, userHandler.GetUserByID)

	r.GET("/site", middleware, siteHandler.GetAllSiteByUser)
	r.POST("/site", middleware, siteHandler.CreateSite)
	r.GET("/site/:site_id", middleware, siteHandler.GetSiteByID)
	r.PUT("/site/:site_id", middleware, siteHandler.UpdateSite)
	r.DELETE("/site/:site_id", middleware, siteHandler.DeleteSite)

	r.Run()
}
