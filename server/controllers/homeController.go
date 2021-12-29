package controllers

import (
	"tripit/utils"

	"github.com/gin-gonic/gin"
)

func GetHome() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, utils.ResponseMessage{
			Success: true,
			Message: "Home Page",
		})
	}
}
