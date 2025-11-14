package response

type Post_Api_FishingSpot_Mark_Response struct {
	// 钓鱼点标记ID
	FishingSpotMarkerID uint `json:"fishing_spot_marker_id"`
}

type Get_Api_FishingSpot_RectList_Response struct {
	// 钓鱼点标记ID
	FishingSpotMarkerID uint `json:"fishing_spot_marker_id"`

	// 经度
	Longitude float64 `json:"longitude"`
	// 纬度
	Latitude float64 `json:"latitude"`

	// 是否允许车辆通行
	CarAccess bool `json:"car_access"`
	// 位置名
	LocationName string `json:"location_name"`

	// 创建者ID
	CreatorID uint `json:"creator_id"`
	// 创建者昵称
	CreatorNickName string `json:"creator_nick_name"`

	// 分类ID列表
	CategoryIDs []uint `json:"category_ids"`
}
