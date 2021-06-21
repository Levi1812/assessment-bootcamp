package handler

import (
	"assessment/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service}
}

// Handler

func (h *userHandler) GetAllUser(c *gin.Context) {
	users, err := h.service.GetAllUser()

	if err != nil {
		c.JSON(500, gin.H{"ERROR": err.Error()})
		return
	}
	c.JSON(201, users)
}

func (h *userHandler) GetUserByID(c *gin.Context) {
	id := c.Params.ByName("user_id")

	user, err := h.service.GetUserByID(id)

	if err != nil {
		c.JSON(500, gin.H{"ERROR": err.Error()})
		return
	}
	c.JSON(201, user)
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.UserRegisterinput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"ERROR": err.Error()})
		return
	}
	user, err := h.service.RegisterUser(input)

	if err != nil {
		c.JSON(500, gin.H{"ERROR": err.Error()})
		return
	}
	c.JSON(201, user)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input user.UserLogininput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"ERROR": err.Error()})
		return
	}
	user, err := h.service.LoginUser(input)

	if err != nil {
		c.JSON(500, gin.H{"ERROR": err.Error()})
		return
	}
	c.JSON(201, user)
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	UserID := c.Params.ByName("book_id")

	var input book.UpdateUser
}
