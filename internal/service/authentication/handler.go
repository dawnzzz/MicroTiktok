package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/dawnzzz/MicroTiktok/internal/service/authentication/auth"
	authentication "github.com/dawnzzz/MicroTiktok/kitex_gen/authentication"
	"github.com/dawnzzz/MicroTiktok/kitex_gen/base"
	"github.com/dawnzzz/MicroTiktok/pkg/e"
)

// AuthenticationServiceImpl implements the last service interface defined in the IDL.
type AuthenticationServiceImpl struct{}

// GenerateToken implements the AuthenticationServiceImpl interface.
func (s *AuthenticationServiceImpl) GenerateToken(ctx context.Context, req *authentication.GenerateTokenRequest) (resp *authentication.GenerateTokenResponse, err error) {
	// TODO: Your code here...
	return
}

// CheckToken implements the AuthenticationServiceImpl interface.
func (s *AuthenticationServiceImpl) CheckToken(ctx context.Context, req *authentication.CheckTokenRequest) (resp *authentication.CheckTokenResponse, err error) {
	// 解析token是否正确
	userClaims, err := auth.ParseToken(req.Token)
	if err != nil { // 不正确
		return nil, kerrors.NewBizStatusError(e.ErrAuthenticationFailed, e.GetErrMsg(e.ErrAuthenticationFailed))
	}

	// 正确，返回user id
	return &authentication.CheckTokenResponse{
		BaseResp: &base.BaseResponse{
			StatusCode: e.Success,
			StatusMsg:  e.GetErrMsg(e.Success),
		},
		UserId: userClaims.UserID,
	}, nil
}
