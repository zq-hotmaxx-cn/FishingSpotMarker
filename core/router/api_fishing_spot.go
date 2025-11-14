package router

import (
	"FishingSpotMarker/core/controller"
	"FishingSpotMarker/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// 路由: /api/fishing_spot
func loadApiFishingSpotRouter(r *gin.RouterGroup) {
	fishing_spot_router := r.Group("/fishing_spot")
	fishing_spot_router.Use(middleware.Auth())
	{
		// 标记钓点
		fishing_spot_router.POST("/mark", controller.Post_Api_FishingSpot_Mark)
		// 获取指定矩形范围钓点列表
		fishing_spot_router.GET("/rect-list", controller.Get_Api_FishingSpot_RectList)
		// 获取指定地区钓点列表
		fishing_spot_router.GET("/area-list", controller.Get_Default)
		// 获取指定钓点详情
		fishing_spot_router.GET("/detail")
		// 获取多个钓点详情
		fishing_spot_router.GET("/details")
	}
}
