package repository

import (
	"FishingSpotMarker/core/data"
	"FishingSpotMarker/pkg/model"
)

// 创建分类
func CreateCategory(category *model.Category) error {
	return data.GetDataBase().Create(category).Error
}

// 根据分类ID列表获取已存在的分类列表
func FindCategoryListByIDs(ids []uint) ([]model.Category, error) {
	var categories []model.Category
	err := data.GetDataBase().Where("id IN ?", ids).Find(&categories).Error
	return categories, err
}
