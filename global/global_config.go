package global

import "time"

var ConfigObj *Config

// Config 全局配置
type Config struct {
	MysqlConfig
}

// MysqlConfig Mysql配置
type MysqlConfig struct {
	Host            string // 地址
	Port            int    // 端口
	User            string // 用户名
	Password        string // 密码
	DBName          string // 数据库名
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}
