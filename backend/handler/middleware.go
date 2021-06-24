package handler

import (
	"assessment/auth"
	"assessment/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Middleware(userService user.Service, authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		if header == "" || len(header) == 0 {

			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorize user"})
			return
		}

		token, err := authService.ValidateToken(header)

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorize user"})
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorize user"})
			return
		}

		userID := int(claim["user_id"].(float64))

		c.Set("CurrentUser", userID)
	}
}
