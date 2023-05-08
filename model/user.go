package model

import "gorm.io/gorm"

// User 用户表
type User struct {
	gorm.Model
	UserID          int64  `gorm:"unique;not null"` // 用户ID
	Name            string // 用户名
	Avatar          string // 头像
	BackgroundImage string // 背景图片
	Signature       string // 用户签名
}
