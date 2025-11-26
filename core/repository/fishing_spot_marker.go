package repository

import (
	"FishingSpotMarker/core/data"
	"FishingSpotMarker/pkg/model"

	"gorm.io/gorm"
)

// 使用事务创建钓点标记
func CreateFishingSpotMarkerWithTransaction(
	t *gorm.DB,
	fishing_spot_marker *model.FishingSpotMarker,
) error {
	return t.Create(fishing_spot_marker).Error
}

// 通过经纬度矩形获取范围内所有的钓点标记
func FindFishingSpotMarkerListByRect(
	min_latitude float64,
	min_longitude float64,
	max_latitude float64,
	max_longitude float64,
) ([]model.FishingSpotMarker, error) {
	var fishing_spot_marker_list []model.FishingSpotMarker
	var user *model.User
	var category *model.Category
	err := data.GetDataBase().
		Where(
			"latitude >= ? AND latitude <= ? AND longitude >= ? AND longitude <= ?",
			min_latitude,
			max_latitude,
			min_longitude,
			max_longitude,
		).
		Preload(user.TableName()).
		Preload(category.TableName()).
		Find(&fishing_spot_marker_list).Error
	return fishing_spot_marker_list, err
}

func FirstFishingSpotMarkerWithID(id uint) (*model.FishingSpotMarker, error) {
	fishing_spot_marker := &model.FishingSpotMarker{
		Model: gorm.Model{
			ID: id,
		},
	}
	var user *model.User
	var category *model.Category
	err := data.GetDataBase().
		Preload(user.TableName()).
		Preload(category.TableName()).
		First(fishing_spot_marker).Error
	return fishing_spot_marker, err
}
