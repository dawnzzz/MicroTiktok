package main

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/dawnzzz/MicroTiktok/internal/service/id_generator/config"
	"github.com/dawnzzz/MicroTiktok/kitex_gen/base"
	id_generator "github.com/dawnzzz/MicroTiktok/kitex_gen/id_generator"
	"github.com/dawnzzz/MicroTiktok/pkg/e"
	"sync"
)

// IdGeneratorServiceImpl implements the last service interface defined in the IDL.
type IdGeneratorServiceImpl struct {
	UserIDGenerator    *snowflake.Node
	VideoIDGenerator   *snowflake.Node
	CommentIDGenerator *snowflake.Node
	mutex              sync.Mutex
}

// GetUserID implements the IdGeneratorServiceImpl interface.
func (s *IdGeneratorServiceImpl) GetUserID(ctx context.Context, req *id_generator.GetUserIdRequest) (resp *id_generator.GetUserIdResponse, err error) {
	if s.UserIDGenerator == nil {
		// 饿汉式初始化
		s.mutex.Lock()
		defer s.mutex.Unlock()
		if s.UserIDGenerator == nil {
			s.UserIDGenerator, err = snowflake.NewNode(config.IdGeneratorConfigObj.NodeID)
			if err != nil {
				return nil, kerrors.NewBizStatusError(e.ErrUserIDGeneratorInitFailed, e.GetErrMsg(e.ErrUserIDGeneratorInitFailed))
			}
		}
	}

	return &id_generator.GetUserIdResponse{
		BaseResp: &base.BaseResponse{
			StatusCode: e.Success,
			StatusMsg:  e.GetErrMsg(e.Success),
		},
		UserId: s.UserIDGenerator.Generate().Int64(),
	}, nil
}

// GetVideoID implements the IdGeneratorServiceImpl interface.
func (s *IdGeneratorServiceImpl) GetVideoID(ctx context.Context, req *id_generator.GetVideoIdRequest) (resp *id_generator.GetVideoIdResponse, err error) {
	if s.VideoIDGenerator == nil {
		// 饿汉式初始化
		s.mutex.Lock()
		defer s.mutex.Unlock()
		if s.VideoIDGenerator == nil {
			s.VideoIDGenerator, err = snowflake.NewNode(config.IdGeneratorConfigObj.NodeID)
			if err != nil {
				return nil, kerrors.NewBizStatusError(e.ErrVideoIDGeneratorInitFailed, e.GetErrMsg(e.ErrVideoIDGeneratorInitFailed))
			}
		}
	}

	return &id_generator.GetVideoIdResponse{
		BaseResp: &base.BaseResponse{
			StatusCode: e.Success,
			StatusMsg:  e.GetErrMsg(e.Success),
		},
		VideoId: s.VideoIDGenerator.Generate().Int64(),
	}, nil
}

// GetCommentID implements the IdGeneratorServiceImpl interface.
func (s *IdGeneratorServiceImpl) GetCommentID(ctx context.Context, req *id_generator.GetCommentIdRequest) (resp *id_generator.GetCommentIdResponse, err error) {
	if s.CommentIDGenerator == nil {
		// 饿汉式初始化
		s.mutex.Lock()
		defer s.mutex.Unlock()
		if s.CommentIDGenerator == nil {
			s.CommentIDGenerator, err = snowflake.NewNode(config.IdGeneratorConfigObj.NodeID)
			if err != nil {
				return nil, kerrors.NewBizStatusError(e.ErrCommentIDGeneratorInitFailed, e.GetErrMsg(e.ErrCommentIDGeneratorInitFailed))
			}
		}
	}

	return &id_generator.GetCommentIdResponse{
		BaseResp: &base.BaseResponse{
			StatusCode: e.Success,
			StatusMsg:  e.GetErrMsg(e.Success),
		},
		CommentId: s.CommentIDGenerator.Generate().Int64(),
	}, nil
}
