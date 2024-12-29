package middlewares

import (
	"net/http"
	"story-plateform/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized request",
			})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized request",
			})
			c.Abort()
			return
		}

		c.Set("username", claims["username"])
		c.Set("email", claims["email"])
		c.Next()
	}
}
