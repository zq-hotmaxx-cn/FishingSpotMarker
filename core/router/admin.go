package router

import "FishingSpotMarker/core/controller"

func loadAdminRouter() {
	admin_router := fishing_spot_marker_router.Group("/admin")
	{
		// 健康检查
		admin_router.GET("/health", controller.Get_Admin_Health)
	}
}
