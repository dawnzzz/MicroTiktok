package main

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/dawnzzz/MicroTiktok/global"
	"github.com/dawnzzz/MicroTiktok/kitex_gen/base"
	id_generator "github.com/dawnzzz/MicroTiktok/kitex_gen/id_generator"
	"github.com/dawnzzz/MicroTiktok/pkg/e"
	"strconv"
	"sync"
	"time"
)

// IdGeneratorServiceImpl implements the last service interface defined in the IDL.
type IdGeneratorServiceImpl struct {
	UserIDGenerator    *snowflake.Node
	VideoIDGenerator   *snowflake.Node
	CommentIDGenerator *snowflake.Node
	mutex              sync.Mutex
}

const (
	userField    = "user"
	videoField   = "video"
	commentField = "comment"
)

// GetUserID implements the IdGeneratorServiceImpl interface.
func (s *IdGeneratorServiceImpl) GetUserID(ctx context.Context, req *id_generator.GetUserIdRequest) (resp *id_generator.GetUserIdResponse, err error) {
	if s.UserIDGenerator == nil {
		// 饿汉式初始化
		s.mutex.Lock()
		if s.UserIDGenerator == nil {
			s.UserIDGenerator, err = s.NewIdGenerator(userField)
			if err != nil {
				return nil, kerrors.NewBizStatusError(e.ErrUserIDGeneratorInitFailed, e.GetErrMsg(e.ErrUserIDGeneratorInitFailed))
			}
		}
		s.mutex.Unlock()
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
		if s.VideoIDGenerator == nil {
			s.VideoIDGenerator, err = s.NewIdGenerator(videoField)
			if err != nil {
				return nil, kerrors.NewBizStatusError(e.ErrVideoIDGeneratorInitFailed, e.GetErrMsg(e.ErrVideoIDGeneratorInitFailed))
			}
		}
		s.mutex.Unlock()
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
		if s.CommentIDGenerator == nil {
			s.CommentIDGenerator, err = s.NewIdGenerator(commentField)
			if err != nil {
				return nil, kerrors.NewBizStatusError(e.ErrCommentIDGeneratorInitFailed, e.GetErrMsg(e.ErrCommentIDGeneratorInitFailed))
			}
		}
		s.mutex.Unlock()
	}

	return &id_generator.GetCommentIdResponse{
		BaseResp: &base.BaseResponse{
			StatusCode: e.Success,
			StatusMsg:  e.GetErrMsg(e.Success),
		},
		CommentId: s.CommentIDGenerator.Generate().Int64(),
	}, nil
}

const (
	curNodeIdKeyFormat     = "microTiktok/id_generator/%v/curNodeId"
	curNodeIdKeyLockFormat = "microTiktok/id_generator/%v/curNodeIdLock"
)

func (s *IdGeneratorServiceImpl) NewIdGenerator(field string) (*snowflake.Node, error) {
	c, err := clientv3.New(clientv3.Config{
		Endpoints: global.EtcdEndpoints,
	})
	if err != nil {
		return nil, err
	}
	defer c.Close()

	session, err := concurrency.NewSession(c)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	// 获取分布式锁
	curNodeIdKeyLock := fmt.Sprintf(curNodeIdKeyLockFormat, field)
	m := concurrency.NewMutex(session, curNodeIdKeyLock)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err = m.Lock(ctx)
	defer cancel()
	if err != nil {
		return nil, err
	}
	defer m.Unlock(ctx)

	// 获取当前节点的node id
	curNodeIdKey := fmt.Sprintf(curNodeIdKeyFormat, field)
	getCtx, getCancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := c.Get(getCtx, curNodeIdKey)
	defer getCancel()
	if err != nil {
		return nil, err
	}

	var id int64
	if resp.Count != 0 {
		id, err = strconv.ParseInt(string(resp.Kvs[0].Value), 10, 64)
		if err != nil {
			return nil, err
		}
	}

	// 创建node
	generator, err := snowflake.NewNode(id)
	if err != nil {
		return nil, err
	}

	// id++，并且写回etcd
	id++
	putCtx, putCancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = c.Put(putCtx, curNodeIdKey, strconv.FormatInt(id, 10))
	defer putCancel()
	if err != nil {
		return nil, err
	}

	return generator, nil
}
