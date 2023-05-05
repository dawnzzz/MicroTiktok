package main

import (
	"context"
	comment "github.com/dawnzzz/MicroTiktok/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// Comment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) Comment(ctx context.Context, req *comment.CommentRequest) (resp *comment.CommentResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentDelete implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentDelete(ctx context.Context, req *comment.CommentDeleteRequest) (resp *comment.CommentDeleteResponse, err error) {
	// TODO: Your code here...
	return
}

// GetCommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) GetCommentList(ctx context.Context, req *comment.GetCommentListRequest) (resp *comment.GetCommentListResponse, err error) {
	// TODO: Your code here...
	return
}
