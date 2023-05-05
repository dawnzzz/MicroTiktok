package main

import (
	"context"
	favorite "github.com/dawnzzz/MicroTiktok/kitex_gen/favorite"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// Favorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) Favorite(ctx context.Context, req *favorite.FavoriteRequest) (resp *favorite.FavoriteResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteCancel implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteCancel(ctx context.Context, req *favorite.FavoriteCancelRequest) (resp *favorite.FavoriteCancelResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFavoriteVideoList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) GetFavoriteVideoList(ctx context.Context, req *favorite.GetFavoriteVideoListRequest) (resp *favorite.GetFavoriteVideoListResponse, err error) {
	// TODO: Your code here...
	return
}
