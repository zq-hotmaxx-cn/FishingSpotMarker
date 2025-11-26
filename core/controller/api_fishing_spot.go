package controller

import (
	"FishingSpotMarker/core/service"
	"FishingSpotMarker/core/validate"
	"FishingSpotMarker/pkg/http/request"
	"FishingSpotMarker/pkg/utils/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Post_Api_FishingSpot_Mark(c *gin.Context) {
	// 绑定请求参数
	req := &request.Post_Api_FishingSpot_Mark_Request{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	// 验证请求参数
	if err := validate.Validate(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 获取当前用户
	auth_user, exists := user.GetAuthUserWithGinContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	// 调用服务
	resp, err := service.Post_Api_FishingSpot_Mark_Service(req, auth_user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, resp)
}

func Get_Api_FishingSpot_RectList(c *gin.Context) {
	// 绑定请求参数
	req := &request.Get_Api_FishingSpot_RectList_Request{}
	if err := req.BindQueryWithGinContext(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	// 验证请求参数
	if err := validate.Validate(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 获取当前用户
	auth_user, exists := user.GetAuthUserWithGinContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	// 调用服务
	resp, err := service.Get_Api_FishingSpot_RectList_Service(req, auth_user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, resp)
}

func Get_Api_FishingSpot_Detail(c *gin.Context) {
	// 绑定请求参数
	req := &request.Get_Api_FishingSpot_Detail_Request{}
	if err := req.BindQueryWithGinContext(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	// 验证请求参数
	if err := validate.Validate(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 获取当前用户
	auth_user, exists := user.GetAuthUserWithGinContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	// 调用服务
	resp, err := service.Get_Api_FishingSpot_Detail_Service(req, auth_user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, resp)
}
