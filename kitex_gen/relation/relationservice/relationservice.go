// Code generated by Kitex v0.5.2. DO NOT EDIT.

package relationservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	relation "github.com/dawnzzz/MicroTiktok/kitex_gen/relation"
)

func serviceInfo() *kitex.ServiceInfo {
	return relationServiceServiceInfo
}

var relationServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "RelationService"
	handlerType := (*relation.RelationService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Relation":        kitex.NewMethodInfo(relationHandler, newRelationServiceRelationArgs, newRelationServiceRelationResult, false),
		"RelationCancel":  kitex.NewMethodInfo(relationCancelHandler, newRelationServiceRelationCancelArgs, newRelationServiceRelationCancelResult, false),
		"GetFollowList":   kitex.NewMethodInfo(getFollowListHandler, newRelationServiceGetFollowListArgs, newRelationServiceGetFollowListResult, false),
		"GetFollowerList": kitex.NewMethodInfo(getFollowerListHandler, newRelationServiceGetFollowerListArgs, newRelationServiceGetFollowerListResult, false),
		"GetFriendList":   kitex.NewMethodInfo(getFriendListHandler, newRelationServiceGetFriendListArgs, newRelationServiceGetFriendListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "relation",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func relationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceRelationArgs)
	realResult := result.(*relation.RelationServiceRelationResult)
	success, err := handler.(relation.RelationService).Relation(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceRelationArgs() interface{} {
	return relation.NewRelationServiceRelationArgs()
}

func newRelationServiceRelationResult() interface{} {
	return relation.NewRelationServiceRelationResult()
}

func relationCancelHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceRelationCancelArgs)
	realResult := result.(*relation.RelationServiceRelationCancelResult)
	success, err := handler.(relation.RelationService).RelationCancel(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceRelationCancelArgs() interface{} {
	return relation.NewRelationServiceRelationCancelArgs()
}

func newRelationServiceRelationCancelResult() interface{} {
	return relation.NewRelationServiceRelationCancelResult()
}

func getFollowListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceGetFollowListArgs)
	realResult := result.(*relation.RelationServiceGetFollowListResult)
	success, err := handler.(relation.RelationService).GetFollowList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceGetFollowListArgs() interface{} {
	return relation.NewRelationServiceGetFollowListArgs()
}

func newRelationServiceGetFollowListResult() interface{} {
	return relation.NewRelationServiceGetFollowListResult()
}

func getFollowerListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceGetFollowerListArgs)
	realResult := result.(*relation.RelationServiceGetFollowerListResult)
	success, err := handler.(relation.RelationService).GetFollowerList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceGetFollowerListArgs() interface{} {
	return relation.NewRelationServiceGetFollowerListArgs()
}

func newRelationServiceGetFollowerListResult() interface{} {
	return relation.NewRelationServiceGetFollowerListResult()
}

func getFriendListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceGetFriendListArgs)
	realResult := result.(*relation.RelationServiceGetFriendListResult)
	success, err := handler.(relation.RelationService).GetFriendList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceGetFriendListArgs() interface{} {
	return relation.NewRelationServiceGetFriendListArgs()
}

func newRelationServiceGetFriendListResult() interface{} {
	return relation.NewRelationServiceGetFriendListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Relation(ctx context.Context, req *relation.RelationRequest) (r *relation.RelationResponse, err error) {
	var _args relation.RelationServiceRelationArgs
	_args.Req = req
	var _result relation.RelationServiceRelationResult
	if err = p.c.Call(ctx, "Relation", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RelationCancel(ctx context.Context, req *relation.RelationCancelRequest) (r *relation.RelationCancelResponse, err error) {
	var _args relation.RelationServiceRelationCancelArgs
	_args.Req = req
	var _result relation.RelationServiceRelationCancelResult
	if err = p.c.Call(ctx, "RelationCancel", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFollowList(ctx context.Context, req *relation.GetRelationFollowListRequest) (r *relation.GetRelationFollowListResponse, err error) {
	var _args relation.RelationServiceGetFollowListArgs
	_args.Req = req
	var _result relation.RelationServiceGetFollowListResult
	if err = p.c.Call(ctx, "GetFollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFollowerList(ctx context.Context, req *relation.GetRelationFollowerListRequest) (r *relation.GetRelationFollowerListResponse, err error) {
	var _args relation.RelationServiceGetFollowerListArgs
	_args.Req = req
	var _result relation.RelationServiceGetFollowerListResult
	if err = p.c.Call(ctx, "GetFollowerList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFriendList(ctx context.Context, req *relation.GetRelationFriendListRequest) (r *relation.GetRelationFriendListResponse, err error) {
	var _args relation.RelationServiceGetFriendListArgs
	_args.Req = req
	var _result relation.RelationServiceGetFriendListResult
	if err = p.c.Call(ctx, "GetFriendList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
