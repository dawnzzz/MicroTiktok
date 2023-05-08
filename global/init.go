package global

import "time"

func init() {
	ConfigObj = &Config{
		MysqlConfig: MysqlConfig{
			Host:            "127.0.0.1",
			Port:            3306,
			User:            "root",
			Password:        "123456",
			DBName:          "micro_tiktok",
			MaxIdleConns:    10,
			MaxOpenConns:    100,
			ConnMaxLifetime: time.Hour,
		},
	}
}