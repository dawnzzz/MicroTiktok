package main

import (
	"context"
	relation "github.com/dawnzzz/MicroTiktok/kitex_gen/relation"
)

// SocialityServiceImpl implements the last service interface defined in the IDL.
type SocialityServiceImpl struct{}

// Relation implements the SocialityServiceImpl interface.
func (s *SocialityServiceImpl) Relation(ctx context.Context, req *relation.RelationRequest) (resp *relation.RelationResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationCancel implements the SocialityServiceImpl interface.
func (s *SocialityServiceImpl) RelationCancel(ctx context.Context, req *relation.RelationCancelRequest) (resp *relation.RelationCancelResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowList implements the SocialityServiceImpl interface.
func (s *SocialityServiceImpl) GetFollowList(ctx context.Context, req *relation.GetRelationFollowListRequest) (resp *relation.GetRelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowerList implements the SocialityServiceImpl interface.
func (s *SocialityServiceImpl) GetFollowerList(ctx context.Context, req *relation.GetRelationFollowerListRequest) (resp *relation.GetRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFriendList implements the SocialityServiceImpl interface.
func (s *SocialityServiceImpl) GetFriendList(ctx context.Context, req *relation.GetRelationFriendListRequest) (resp *relation.GetRelationFriendListResponse, err error) {
	// TODO: Your code here...
	return
}
