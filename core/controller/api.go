package controller

import "github.com/gin-gonic/gin"

func Get_Api_Health(c *gin.Context) {
	c.JSON(200, nil)
}
