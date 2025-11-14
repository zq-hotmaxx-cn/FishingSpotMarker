package middleware

import (
	"FishingSpotMarker/pkg/utils/user"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		authorization_slice := strings.Split(authorization, " ")
		if len(authorization_slice) != 2 || authorization_slice[0] != "Bearer" {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		token := authorization_slice[1]

		auth_user, err := auth(token)
		if err != nil {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		user.SetAuthUserWithGinContext(c, auth_user)
	}
}

func auth(token string) (user.User, error) {
	if token == "" {
		return user.User{
			ID:       0,
			Platform: user.PLATFORM_UNKNOWN,
		}, nil
	}

	return user.User{
		ID:       0,
		Platform: user.PLATFORM_UNKNOWN,
	}, nil
}
