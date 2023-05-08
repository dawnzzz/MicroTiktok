package e

import (
	apibase "github.com/dawnzzz/MicroTiktok/internal/api/biz/model/base"
	rpcbase "github.com/dawnzzz/MicroTiktok/kitex_gen/base"
)

var errMsg = map[int32]string{
	Success:                 "Success",
	ErrUnknown:              "未知的错误",
	ErrBadRequest:           "请求参数错误",
	ErrAuthenticationFailed: "用户身份验证失败，请重新登录",

	ErrUserRpcFailed: "User RPC 请求失败",
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
