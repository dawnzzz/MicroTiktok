package model

import "gorm.io/gorm"

// User 用户表
type User struct {
	gorm.Model
	UserID          int64  `gorm:"unique;not null"` // 用户ID
	Name            string `gorm:"unique;not null"` // 用户名
	Password        string // 密码
	Avatar          string // 头像
	BackgroundImage string // 背景图片
	Signature       string // 用户签名
}

// Add 添加一个用户
func (user *User) Add() error {
	return db.Model(user).Create(user).Error
}

// FindOneByUserId 根据user id查询一个用户
func (user *User) FindOneByUserId(userID int64) error {
	return user.FindOne("user_id = ?", userID)
}

// FindOneByPasswordAndUserName 根据用户名和密码查询一个用户
func (user *User) FindOneByPasswordAndUserName(name, password string) error {
	return user.FindOne("name = ? and password = ?", name, password)
}

// FindOne 根据条件查询一个用户
func (user *User) FindOne(query string, args ...interface{}) error {
	return db.Model(user).Where(query, args...).First(user).Error
}
