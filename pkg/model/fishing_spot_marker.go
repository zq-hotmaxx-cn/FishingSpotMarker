package model

import "gorm.io/gorm"

// 钓点标记
type FishingSpotMarker struct {
	gorm.Model

	// 创建者ID
	CreatorID uint

	// 经度
	Longitude float64 `gorm:"type:decimal(10,7); not null" validate:"required, min=-180, max=180"`
	// 纬度
	Latitude float64 `gorm:"type:decimal(9,7); not null" validate:"required, min=-90, max=90"`

	// 是否允许车辆
	CarAccess bool

	// 位置名称
	LocationName string `gorm:"type:varchar(128); not null" validate:"min=1, max=128"`
	// 位置
	Location string `gorm:"type:varchar(256); not null" validate:"min=1, max=128"`
	// 描述
	Description string `gorm:"type:varchar(1024); not null" validate:"min=0, max=1024"`

	// 创建者
	User *User `gorm:"foreignKey:UserID"`
	// 分类列表
	Categories []Category `gorm:"many2many:fishing_spot_marker_categories;"`
}

func (f *FishingSpotMarker) TableName() string {
	return "fishing_spot_markers"
}
