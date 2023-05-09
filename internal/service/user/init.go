package main

import (
	"github.com/dawnzzz/MicroTiktok/global"
	"github.com/dawnzzz/MicroTiktok/initialize"
	"github.com/dawnzzz/MicroTiktok/internal/service/user/config"
	"github.com/dawnzzz/MicroTiktok/model"
)

func InitUserConfig() {
	config.UserConfigObj = &config.UserConfig{
		Port: global.DefaultRpcUserServicePort,
	}
}

// InitDB 初始化数据库
func InitDB() {
	model.InitDB()
}

// InitRpcClient 初始化rpc客户端
func InitRpcClient() {
	// authentication rpc
	initialize.InitRpcAuthenticationClient()

	// society rpc
	initialize.InitRpcSocietyClient()

	// interact rpc
	initialize.InitRpcInteractClient()

	// id generator rpc
	initialize.InitRpcIdGeneratorClient()
}
