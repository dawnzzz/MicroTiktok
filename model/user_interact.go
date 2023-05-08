package model

type UserInteract struct {
	ID             uint  `gorm:"primarykey;autoIncrement"`
	UserID         int64 `gorm:"index"` // 用户ID
	TotalFavorited int64 // 被赞数量
	WorkCount      int64 // 发布作品数量
	FavoriteCount  int64 // 点赞数量
}
