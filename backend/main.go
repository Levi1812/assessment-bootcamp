package main

import (
	"assessment/config"
	"assessment/handler"
	"assessment/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB = config.Config()

	userRepository = user.NewRepo(DB)
	userService    = user.NewService(userRepository)
	userHandler    = handler.NewUserHandler(userService)
)

func Main() {
	r := gin.Default()

	r.GET("/users", userHandler.GetAllUser)
	r.POST("/users/register", userHandler.RegisterUser)
	r.POST("/users/login", userHandler.LoginUser)
	r.GET("users/:user_id", userHandler.GetUserByID)
	r.PUT("users/:user_id")
	r.DELETE("users/:user_id")

	r.Run()
}
