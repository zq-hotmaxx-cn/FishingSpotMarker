package router

import (
	"FishingSpotMarker/core/controller"
	"FishingSpotMarker/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// 路由: /api/category
func loadApiCategoryRouter(r *gin.RouterGroup) {
	category_router := r.Group("/category")
	category_router.Use(middleware.Auth())
	{
		// 创建分类
		category_router.POST("/", controller.Post_Api_Category)
		// 获取所有分类列表
		category_router.GET("/list", controller.Get_Api_Category_List)
		// 获取分类详情
		category_router.GET("/detail", controller.Get_Api_Category_Detail)
	}
}
