package main

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/dawnzzz/MicroTiktok/global"
	"github.com/dawnzzz/MicroTiktok/kitex_gen/authentication"
	"github.com/dawnzzz/MicroTiktok/kitex_gen/base"
	"github.com/dawnzzz/MicroTiktok/kitex_gen/id_generator"
	user "github.com/dawnzzz/MicroTiktok/kitex_gen/user"
	"github.com/dawnzzz/MicroTiktok/model"
	"github.com/dawnzzz/MicroTiktok/pkg/e"
	"gorm.io/gorm"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	// 获取 user id
	getUserIdResponse, err := global.RpcIdGeneratorClient.GetUserID(ctx, &id_generator.GetUserIdRequest{})
	if err != nil {
		return nil, err
	}

	// 尝试插入数据
	u := &model.User{
		UserID:   getUserIdResponse.UserId,
		Name:     req.Username,
		Password: encryptPassword(req.Username, req.Password),
	}
	err = u.Add() // 插入数据
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) { // 用户名重复
			return nil, kerrors.NewBizStatusError(e.ErrDuplicatedUserName, e.GetErrMsg(e.ErrDuplicatedUserName))
		}

		// 其他错误
		return nil, kerrors.NewBizStatusError(e.ErrDBError, e.GetErrMsg(e.ErrDBError))
	}

	// 插入成功，请求RPC，申请token
	generateTokenResponse, err := global.RpcAuthenticationClient.GenerateToken(ctx, &authentication.GenerateTokenRequest{
		UserId: getUserIdResponse.UserId,
	})
	if err != nil {
		return nil, kerrors.NewBizStatusError(e.ErrAuthenticationRpcFailed, e.GetErrMsg(e.ErrAuthenticationRpcFailed))
	}

	// 申请token成功，返回
	return &user.UserRegisterResponse{
		BaseResp: &base.BaseResponse{
			StatusCode: e.Success,
			StatusMsg:  e.GetErrMsg(e.Success),
		},
		UserId: getUserIdResponse.UserId,
		Token:  generateTokenResponse.Token,
	}, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserRequest) (resp *user.GetUserResponse, err error) {
	// TODO: Your code here...
	return
}

// BatchGetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) BatchGetUserInfo(ctx context.Context, req *user.BatchGetUserRequest) (resp *user.BatchGetUserResonse, err error) {
	// TODO: Your code here...
	return
}

const (
	salt1 = "AF#%%#KLNLKN456^*$"
	salt2 = "S$#KJGSSD&&())%%*F"
)

func encryptPassword(name, password string) string {
	h1 := sha256.New()
	h1.Write([]byte(name))
	hash1 := h1.Sum([]byte(salt1))

	h2 := sha256.New()
	h2.Write(hash1)
	h2.Write([]byte(password))
	finialHash := h2.Sum([]byte(salt2))

	return fmt.Sprintf("%x", finialHash)
}
