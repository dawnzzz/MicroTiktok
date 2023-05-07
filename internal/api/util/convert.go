package util

import (
	apibase "github.com/dawnzzz/MicroTiktok/internal/api/biz/model/base"
	rpcbase "github.com/dawnzzz/MicroTiktok/kitex_gen/base"
)

func RpcUserConvertToApiUser(rpcUser *rpcbase.User) *apibase.User {
	return &apibase.User{
		ID:              rpcUser.Id,
		Name:            rpcUser.Name,
		FollowCount:     rpcUser.FollowCount,
		FollowerCount:   rpcUser.FollowerCount,
		IsFollow:        rpcUser.IsFollow,
		Avatar:          rpcUser.Avatar,
		BackgroundImage: rpcUser.BackgroundImage,
		Signature:       rpcUser.Signature,
		TotalFavorited:  rpcUser.TotalFavorited,
		WorkCount:       rpcUser.WorkCount,
		FavoriteCount:   rpcUser.FavoriteCount,
	}
}
