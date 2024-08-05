package middleware

import (
	"net/http"

	"github.com/ARunni/ToDo-Backend-API/Api-service/pkg/utils/helper"
	"github.com/gin-gonic/gin"
)

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := helper.GetTokenFromHeader(authHeader)
		if tokenString == "" {
			var err error
			tokenString, err = c.Cookie("Authorization")
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}
		userID, userName, err := helper.ExtractUserIDFromToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{
				"Invalid": "Invalid Token",
			})
			c.AbortWithStatus(401)
			return
		}

		c.Set("user_id", userID)
		c.Set("user_name", userName)
		c.Next()
	}
}
