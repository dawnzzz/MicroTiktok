package util

import (
	apibase "github.com/dawnzzz/MicroTiktok/internal/api/biz/model/base"
	rpcbase "github.com/dawnzzz/MicroTiktok/kitex_gen/base"
)

func RpcUserConvertToApiUser(rpcUser *rpcbase.User) *apibase.User {
	return &apibase.User{
		ID:   rpcUser.Id,
		Name: rpcUser.Name,
		SocialInfo: &apibase.SocialInfo{
			FollowCount:   rpcUser.SocialInfo.FollowCount,
			FollowerCount: rpcUser.SocialInfo.FollowerCount,
			IsFollow:      rpcUser.SocialInfo.IsFollow,
		},
		Avatar:          rpcUser.Avatar,
		BackgroundImage: rpcUser.BackgroundImage,
		Signature:       rpcUser.Signature,
		UserInteracInfo: &apibase.UserInteractInfo{
			TotalFavorited: rpcUser.UserInteracInfo.TotalFavorited,
			WorkCount:      rpcUser.UserInteracInfo.WorkCount,
			FavoriteCount:  rpcUser.UserInteracInfo.FavoriteCount,
		},
	}
}
