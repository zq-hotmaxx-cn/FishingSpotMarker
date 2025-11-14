package request

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
