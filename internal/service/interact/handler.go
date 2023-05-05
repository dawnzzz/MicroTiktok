package main

import (
	"context"
	interact "github.com/dawnzzz/MicroTiktok/kitex_gen/interact"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}

// GetVideoInteractInfo implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) GetVideoInteractInfo(ctx context.Context, req *interact.GetVideoInteractInfoRequest) (resp *interact.GetVideoInteractInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// BatchGetVideoInteractInfo implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) BatchGetVideoInteractInfo(ctx context.Context, req *interact.BatchGetVideoInteractInfoRequest) (resp *interact.BatchGetVideoInteractInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserInteractInfo implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) GetUserInteractInfo(ctx context.Context, req *interact.GetUserInteractInfoRequest) (resp *interact.GetUserInteractInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// BatchGetUserInteractInfo implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) BatchGetUserInteractInfo(ctx context.Context, req *interact.BatchGetUserInteractInfoRequest) (resp *interact.BatchGetUserInteractInfoResponse, err error) {
	// TODO: Your code here...
	return
}
