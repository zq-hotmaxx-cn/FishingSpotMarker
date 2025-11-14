package router

import (
	"FishingSpotMarker/core/controller"
)

// 路由: /api
func loadApiRouter() {
	api_router := fishing_spot_marker_router.Group("/api")
	{
		// 健康检查
		api_router.GET("/health", controller.Get_Api_Health)

		// 发送验证码
		api_router.POST("/captcha", controller.Get_Default)

		// Web服务
		loadApiWebRouter(api_router)

		// 微信服务
		loadApiWechatRouter(api_router)

		// 用户服务
		loadApiUserRouter(api_router)

		// 分类服务
		loadApiCategoryRouter(api_router)

		// 钓点服务
		loadApiFishingSpotRouter(api_router)
	}
}
