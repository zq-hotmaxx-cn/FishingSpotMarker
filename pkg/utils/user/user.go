package user

import (
	"github.com/gin-gonic/gin"
)

type Platform uint8

const (
	PLATFORM_UNKNOWN Platform = iota
	PLATFORM_WEB
	PLATFORM_WECHAT
)

type User struct {
	ID       uint
	Platform Platform
}

func SetAuthUserWithGinContext(c *gin.Context, auth_user User) {
	c.Set("auth_user", auth_user)
}

func GetAuthUserWithGinContext(c *gin.Context) (User, bool) {
	auth_user, exists := c.Get("auth_user")
	return auth_user.(User), exists
}
