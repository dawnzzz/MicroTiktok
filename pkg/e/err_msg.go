package e

import (
	apibase "github.com/dawnzzz/MicroTiktok/internal/api/biz/model/base"
	rpcbase "github.com/dawnzzz/MicroTiktok/kitex_gen/base"
)

var errMsg = map[int32]string{
	Success:                    "Success",
	ErrUnknown:                 "未知的错误",
	ErrBadRequest:              "请求参数错误",
	ErrAuthenticationFailed:    "用户身份验证失败，请重新登录",
	ErrUserNameOrPasswordWrong: "用户名或者密码错误，请重新输入",
	ErrGenerateTokenFailed:     "生成token失败，请重试",
	ErrDBError:                 "数据库错误，请稍后重试",
	ErrUserRpcFailed:           "User RPC 请求失败",
	ErrAuthenticationRpcFailed: "Authentication RPC请求失败",
	ErrDuplicatedUserName:      "用户名重复",

	ErrUserIDGeneratorInitFailed:    "user id 生成器初始化失败",
	ErrVideoIDGeneratorInitFailed:   "video id 生成器初始化失败",
	ErrCommentIDGeneratorInitFailed: "comment id 生成器初始化失败",
}

func GetErrMsg(errno int32) string {
	if msg, ok := errMsg[errno]; ok {
		return msg
	}

	return errMsg[ErrUnknown]
}

func MakeApiBaseResp(errno int32) *apibase.BaseResponse {
	resp := new(apibase.BaseResponse)

	if msg, ok := errMsg[errno]; ok {
		resp.StatusCode = errno
		resp.StatusMsg = msg
	} else {
		resp.StatusCode = ErrUnknown
		resp.StatusMsg = errMsg[ErrUnknown]
	}

	return resp
}

func MakeRpcBaseResp(errno int32) *rpcbase.BaseResponse {
	resp := new(rpcbase.BaseResponse)

	if msg, ok := errMsg[errno]; ok {
		resp.StatusCode = errno
		resp.StatusMsg = msg
	} else {
		resp.StatusCode = ErrUnknown
		resp.StatusMsg = errMsg[ErrUnknown]
	}

	return resp
}
