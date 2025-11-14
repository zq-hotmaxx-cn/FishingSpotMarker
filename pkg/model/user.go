package model

import (
	"time"

	"gorm.io/gorm"
)

// 性别 0-未知 1-男 2-女
type Gender uint8

const (
	GENDER_UNKNOWN Gender = iota
	GENDER_MALE
	GENDER_FEMALE
)

// User 用户主表结构体
type User struct {
	gorm.Model

	// 密码哈希（可为空，为空时无法使用密码登录）
	PasswordHash string `gorm:"type:varchar(256)" validate:"required, max=256"`

	// 微信开放平台唯一标识（跨应用唯一）
	UnionID string `gorm:"type:varchar(64); uniqueIndex; not null" validate:"required, len=64"`
	// 小程序唯一标识
	OpenID string `gorm:"type:varchar(64); uniqueIndex; not null" validate:"required, len=64"`

	// 昵称
	NickName string `gorm:"type:varchar(128); not null" validate:"required, min=1, max=128"`
	// 头像
	AvatarURL string `gorm:"type:varchar(512); not null" validate:"min=1, max=256"`
	// 性别
	Gender Gender `gorm:"not null; default:0" validate:"required, oneof=0 1 2"`
	// 国家
	Country string `gorm:"type:varchar(64); not null" validate:"min=1, max=128"`
	// 省份
	Province string `gorm:"type:varchar(64); not null" validate:"min=1, max=128"`
	// 城市
	City string `gorm:"type:varchar(64); not null" validate:"min=1, max=128"`
	// 语言
	Language string `gorm:"type:varchar(32); not null" validate:"min=1, max=128"`

	// 最后活跃时间
	LastActive time.Time
	// 过期时间
	ExpiresAt time.Time
}

func (u *User) TableName() string {
	return "users"
}
