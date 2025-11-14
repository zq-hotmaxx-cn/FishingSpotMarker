package router

import "github.com/gin-gonic/gin"

// 路由: /api/wechat
func loadApiWechatRouter(r *gin.RouterGroup) {
	wechat_router := r.Group("/wechat")
	{
		// 登录
		wechat_router.POST("/login")
	}
}
