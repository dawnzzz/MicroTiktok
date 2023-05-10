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
	getUserIdResponse, userIdErr := global.RpcIdGeneratorClient.GetUserID(ctx, &id_generator.GetUserIdRequest{})
	if userIdErr != nil {
		return nil, userIdErr
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
		return nil, err
	}

	// 申请token成功，返回
	resp = &user.UserRegisterResponse{
		BaseResp: &base.BaseResponse{
			StatusCode: e.Success,
			StatusMsg:  e.GetErrMsg(e.Success),
		},
		UserId: getUserIdResponse.UserId,
		Token:  generateTokenResponse.Token,
	}

	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// 在数据库中查找用户
	u := &model.User{}
	err = u.FindOneByPasswordAndUserName(req.Username, encryptPassword(req.Username, req.Password))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 找不到用户，用户名密码错误
			return nil, kerrors.NewBizStatusError(e.ErrUserNameOrPasswordWrong, e.GetErrMsg(e.ErrUserNameOrPasswordWrong))
		}

		return nil, kerrors.NewBizStatusError(e.ErrDBError, e.GetErrMsg(e.ErrDBError))
	}

	// 用户名和密码正确，申请token
	generateTokenResponse, err := global.RpcAuthenticationClient.GenerateToken(ctx, &authentication.GenerateTokenRequest{
		UserId: u.UserID,
	})
	if err != nil {
		return nil, err
	}

	// 登录成功，返回
	resp = &user.UserLoginResponse{
		BaseResp: &base.BaseResponse{
			StatusCode: e.Success,
			StatusMsg:  e.GetErrMsg(e.Success),
		},
		UserId: u.UserID,
		Token:  generateTokenResponse.Token,
	}

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
