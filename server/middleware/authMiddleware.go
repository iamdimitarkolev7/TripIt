package middleware

import (
	"fmt"
	"os"
	"strings"
	"tripit/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieHeaderString := c.Request.Header.Get("Cookie")
		tokens := strings.Split(cookieHeaderString, "=")

		if len(tokens) != 2 || tokens[1] == "" {
			c.Abort()
			c.JSON(401, utils.ResponseMessage{
				Success: false,
				Message: "You are not logged in",
			})
			return
		}

		tokenStr := tokens[1]

		token, parseErr := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error")
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if parseErr != nil {
			c.Abort()
			c.JSON(401, utils.ResponseMessage{
				Success: false,
				Message: "Invalid token",
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("id", claims["id"])
			c.Next()
		} else {
			c.Abort()
			c.JSON(401, utils.ResponseMessage{
				Success: false,
				Message: "There was a problem",
			})
		}
	}
}
