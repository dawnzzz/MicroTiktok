package main

import (
	"github.com/dawnzzz/MicroTiktok/initialize"
	"github.com/dawnzzz/MicroTiktok/model"
)

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
}
