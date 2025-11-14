package repository

import (
	"FishingSpotMarker/core/data"
	"FishingSpotMarker/pkg/model"
)

// 查找用户是否存在，只返回布尔值
func ExistsUser(id uint) (bool, error) {
	var count int64
	err := data.GetDataBase().Model(&model.User{}).Where("id = ?", id).Limit(1).Count(&count).Error
	return count > 0, err
}
