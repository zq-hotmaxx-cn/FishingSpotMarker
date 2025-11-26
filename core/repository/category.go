package repository

import (
	"FishingSpotMarker/core/data"
	"FishingSpotMarker/pkg/model"

	"gorm.io/gorm"
)

// 创建分类
func CreateCategory(category *model.Category) error {
	return data.GetDataBase().Create(category).Error
}

func FindAllCategoryList() ([]model.Category, error) {
	var categories []model.Category
	err := data.GetDataBase().Find(&categories).Error
	return categories, err
}

func FirstCategoryWithID(id uint) (*model.Category, error) {
	category := &model.Category{
		Model: gorm.Model{
			ID: id,
		},
	}
	var user *model.User
	err := data.GetDataBase().
		Preload(user.TableName()). // 预加载用户信息
		First(category).Error
	return category, err
}

// 根据分类ID列表获取已存在的分类列表
func FindCategoryListByIDs(ids []uint) ([]model.Category, error) {
	var categories []model.Category
	err := data.GetDataBase().Where("id IN ?", ids).Find(&categories).Error
	return categories, err
}
