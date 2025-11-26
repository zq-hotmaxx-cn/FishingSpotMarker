package request

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Post_Api_FishingSpot_Mark_Request struct {
	// 经度
	Longitude float64 `json:"longitude" validate:"required, min=-180, max=180"`
	// 纬度
	Latitude float64 `json:"latitude" validate:"required, min=-90, max=90"`

	// 是否允许车辆
	CarAccess bool `json:"car_access"`

	// 位置名称
	LocationName string `json:"location_name" validate:"min=1, max=128"`
	// 位置
	Location string `json:"location" validate:"min=1, max=128"`
	// 描述
	Description string `json:"description" validate:"min=0, max=1024"`

	// 分类ID列表
	CategoryIDs []uint `json:"category_ids"`
}

type Get_Api_FishingSpot_RectList_Request struct {
	// 最小经度
	MinLatitude float64 `json:"min_latitude" validate:"required, min=-90, max=90"`
	// 最小纬度
	MinLongitude float64 `json:"min_longitude" validate:"required, min=-180, max=180"`
	// 最大经度
	MaxLatitude float64 `json:"max_latitude" validate:"required, min=-90, max=90"`
	// 最大纬度
	MaxLongitude float64 `json:"max_longitude" validate:"required, min=-180, max=180"`

	// 是否允许车辆
	CarAccess bool `json:"car_access"`
}

func (req Get_Api_FishingSpot_RectList_Request) BindQueryWithGinContext(c *gin.Context) error {
	query_min_latitude := c.Query("min_latitude")
	query_min_longitude := c.Query("min_longitude")
	query_max_latitude := c.Query("max_latitude")
	query_max_longitude := c.Query("max_longitude")
	query_car_access := c.Query("car_access")

	min_latitude, err := strconv.ParseFloat(query_min_latitude, 64)
	if err != nil {
		return err
	}
	min_longitude, err := strconv.ParseFloat(query_min_longitude, 64)
	if err != nil {
		return err
	}
	max_latitude, err := strconv.ParseFloat(query_max_latitude, 64)
	if err != nil {
		return err
	}
	max_longitude, err := strconv.ParseFloat(query_max_longitude, 64)
	if err != nil {
		return err
	}
	car_access, err := strconv.ParseBool(query_car_access)
	if err != nil {
		return err
	}

	req = Get_Api_FishingSpot_RectList_Request{
		MinLatitude:  min_latitude,
		MinLongitude: min_longitude,
		MaxLatitude:  max_latitude,
		MaxLongitude: max_longitude,
		CarAccess:    car_access,
	}

	return nil
}

type Get_Api_FishingSpot_Detail_Request struct {
	FishingSpotMarkerID uint `json:"fishing_spot_marker_id" validate:"required"`
}

func (req Get_Api_FishingSpot_Detail_Request) BindQueryWithGinContext(c *gin.Context) error {
	query_fishing_spot_marker_id := c.Query("fishing_spot_marker_id")
	fishing_spot_marker_id, err := strconv.ParseUint(query_fishing_spot_marker_id, 10, 64)
	if err != nil {
		return err
	}

	req = Get_Api_FishingSpot_Detail_Request{
		FishingSpotMarkerID: uint(fishing_spot_marker_id),
	}

	return nil
}
