package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/vincentJunior1/test-kriya/models"
)

type JwtClaims struct {
	NewData *models.User
	jwt.StandardClaims
}

func JwtAuth(level string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		newData := &JwtClaims{}
		token, err := jwt.ParseWithClaims(tokenString, newData, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unexpected signing method"})
				c.Abort()
			}

			return []byte("secret"), nil
		})

		if token != nil && err == nil {
			if level == "Admin" {
				if newData.NewData.RoleID != 1 {
					c.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error(), "Messages": "Not Allowed"})
					c.Abort()
					return
				}
			}
			c.Next()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error(), "Messages": "Not Authorized"})
			c.Abort()
			return
		}
	}
}
