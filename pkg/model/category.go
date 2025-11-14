package model

import "gorm.io/gorm"

// 分类
type Category struct {
	gorm.Model

	// 创建者ID
	CreatorID uint
	// 分类名称
	Name string `gorm:"type:varchar(128); unique; not null" validate:"min=1, max=128"`
	// 分类描述
	Description string `gorm:"type:varchar(1024); not null" validate:"min=0, max=1024"`

	// 创建者
	Creator *User `gorm:"foreignKey:CreatorID"`

	// 钓点标记列表
	FishingSpotMarkers []FishingSpotMarker `gorm:"many2many:fishing_spot_marker_categories;"`
}

func (c *Category) TableName() string {
	return "categories"
}
