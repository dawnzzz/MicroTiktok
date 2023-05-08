package model

import "gorm.io/gorm"

type SocialInfo struct {
	ID            uint           `gorm:"primarykey;autoIncrement"`
	UserID        int64          `gorm:"index"` // 用户ID
	FollowCount   int64          // 关注数量
	FollowerCount int64          // 粉丝数量
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
