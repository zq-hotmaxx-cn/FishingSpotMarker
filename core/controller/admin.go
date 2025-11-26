package controller

import "github.com/gin-gonic/gin"

func Get_Admin_Health(c *gin.Context) {
	c.JSON(200, nil)
}
