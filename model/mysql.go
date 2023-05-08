package model

import (
	"fmt"
	"github.com/dawnzzz/MicroTiktok/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.ConfigObj.MysqlConfig.User,
		global.ConfigObj.MysqlConfig.Password,
		global.ConfigObj.MysqlConfig.Host,
		global.ConfigObj.MysqlConfig.Port,
		global.ConfigObj.MysqlConfig.DBName,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(global.ConfigObj.MysqlConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.ConfigObj.MysqlConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(global.ConfigObj.MysqlConfig.ConnMaxLifetime)
}
