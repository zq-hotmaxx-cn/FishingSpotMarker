package router

import (
	"FishingSpotMarker/core/controller"

	"github.com/gin-gonic/gin"
)

// 路由: /api/web
func loadApiWebRouter(r *gin.RouterGroup) {
	api_web_router := r.Group("/web")
	{
		// 注册
		api_web_router.POST("/register", controller.Get_Default)
		// 登录
		api_web_router.POST("/login", controller.Get_Default)
		// 登出
		api_web_router.POST("/logout", controller.Get_Default)
		// 忘记密码
		api_web_router.POST("/forget", controller.Get_Default)
	}
}
