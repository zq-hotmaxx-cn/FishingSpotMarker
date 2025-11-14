package router

import (
	"FishingSpotMarker/core/controller"

	"github.com/gin-gonic/gin"
)

// 路由: /api/user
func loadApiUserRouter(r *gin.RouterGroup) {
	user_router := r.Group("/user")
	{
		// 修改密码
		user_router.POST("/password", controller.Get_Default)
		// 获取用户信息
		user_router.GET("/info")
		// 修改用户信息
		user_router.POST("/info", controller.Get_Default)
		// 删除用户
		user_router.DELETE("/")
	}
}
