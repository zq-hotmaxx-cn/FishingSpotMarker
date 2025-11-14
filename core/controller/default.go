package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get_Default(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "The service has not been implemented yet",
	})
}
