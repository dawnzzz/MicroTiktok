package main

import (
	"context"
	society "github.com/dawnzzz/MicroTiktok/kitex_gen/society"
)

// SocietyServiceImpl implements the last service interface defined in the IDL.
type SocietyServiceImpl struct{}

// GetSocialInfo implements the SocietyServiceImpl interface.
func (s *SocietyServiceImpl) GetSocialInfo(ctx context.Context, req *society.GetSocialInfoRequest) (resp *society.GetSocialInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// BatchGetSocialInfo implements the SocietyServiceImpl interface.
func (s *SocietyServiceImpl) BatchGetSocialInfo(ctx context.Context, req *society.BatchGetSocialInfoRequest) (resp *society.BatchGetSocialInfoResponse, err error) {
	// TODO: Your code here...
	return
}
