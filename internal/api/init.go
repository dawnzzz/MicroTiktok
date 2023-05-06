package main

import "github.com/dawnzzz/MicroTiktok/initialize"

func InitRpcClient() {
	// user rpc
	initialize.InitRpcUserClient()

	// video rpc
	initialize.InitRpcVideoClient()

	// favorite rpc
	initialize.InitRpcFavoriteClient()

	// comment rpc
	initialize.InitRpcCommentClient()

	// relation rpc
	initialize.InitRpcRelationClient()

	// message rpc
	initialize.InitRpcMessageClient()
}
