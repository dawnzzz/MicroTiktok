// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/dawnzzz/MicroTiktok/global"
	"github.com/dawnzzz/MicroTiktok/internal/api/util"
	"github.com/dawnzzz/MicroTiktok/kitex_gen/user"
	"github.com/dawnzzz/MicroTiktok/pkg/e"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	api "github.com/dawnzzz/MicroTiktok/internal/api/biz/model/api"
)

// Register 用户注册
// @router /douyin/user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserRegisterRequest

	resp := new(api.UserRegisterReponse)

	// 验证参数
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = e.ErrBadRequest
		resp.StatusMsg = e.GetErrMsg(resp.StatusCode)
		c.JSON(consts.StatusOK, resp)
		return
	}

	// RPC请求
	registerResponse, err := global.RpcUserClient.Register(ctx, &user.UserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		bizErr, isBizErr := kerrors.FromBizStatusError(err)
		if isBizErr {
			// 是业务类型的错误
			resp.StatusCode = bizErr.BizStatusCode()
			resp.StatusMsg = bizErr.BizMessage()
			c.JSON(consts.StatusOK, resp)
		} else {
			// 是RPC错误
			resp.StatusCode = e.ErrUserRpcFailed
			resp.StatusMsg = e.GetErrMsg(resp.StatusCode)
			c.JSON(consts.StatusInternalServerError, resp)
		}
		return
	}

	// 注册成功
	resp.StatusCode = e.Success
	resp.StatusMsg = e.GetErrMsg(resp.StatusCode)
	resp.UserID = registerResponse.UserId
	resp.Token = registerResponse.Token
	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router /douyin/user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserLoginRequest

	resp := new(api.UserLoginResponse)

	// 验证参数
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = e.ErrBadRequest
		resp.StatusMsg = e.GetErrMsg(resp.StatusCode)
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 发起RPC请求
	loginResponse, err := global.RpcUserClient.Login(ctx, &user.UserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		bizErr, isBizErr := kerrors.FromBizStatusError(err)
		if isBizErr {
			// 是业务类型的错误
			resp.StatusCode = bizErr.BizStatusCode()
			resp.StatusMsg = bizErr.BizMessage()
			c.JSON(consts.StatusOK, resp)
		} else {
			// 是RPC错误
			resp.StatusCode = e.ErrUserRpcFailed
			resp.StatusMsg = e.GetErrMsg(resp.StatusCode)
			c.JSON(consts.StatusInternalServerError, resp)
		}
		return
	}

	// 登录成功
	resp.StatusCode = e.Success
	resp.StatusMsg = e.GetErrMsg(resp.StatusCode)
	resp.UserID = loginResponse.UserId
	resp.Token = loginResponse.Token
	c.JSON(consts.StatusOK, resp)
}

// GetUserInfo .
// @router /douyin/user [GET]
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserRequest

	resp := new(api.UserResponse)

	// 验证参数
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = e.ErrBadRequest
		resp.StatusMsg = e.GetErrMsg(resp.StatusCode)
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 从上下文中获取view id（token解析出的id）
	raw, ok := c.Get(global.HzViewerIDKey)
	if !ok {
		resp.StatusCode = e.ErrBadRequest
		resp.StatusMsg = e.GetErrMsg(resp.StatusCode)
		c.JSON(consts.StatusOK, resp)
		return
	}
	viewerID, ok := raw.(int64)
	if !ok {
		resp.StatusCode = e.ErrBadRequest
		resp.StatusMsg = e.GetErrMsg(resp.StatusCode)
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 发起RPC请求
	userResponse, err := global.RpcUserClient.GetUserInfo(ctx, &user.GetUserRequest{
		ViewerId: viewerID,
		OwnerId:  req.UserID,
	})
	if err != nil {
		resp.StatusCode = e.ErrBadRequest
		resp.StatusMsg = e.GetErrMsg(resp.StatusCode)
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 获取用户信息成功
	resp.StatusCode = e.Success
	resp.StatusMsg = e.GetErrMsg(resp.StatusCode)
	resp.User = util.RpcUserConvertToApiUser(userResponse.User)
	c.JSON(consts.StatusOK, resp)
}
