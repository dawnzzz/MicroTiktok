package main

import (
	"context"
	authentication "github.com/dawnzzz/MicroTiktok/kitex_gen/authentication"
)

// AuthenticationServiceImpl implements the last service interface defined in the IDL.
type AuthenticationServiceImpl struct{}

// GenerateToken implements the AuthenticationServiceImpl interface.
func (s *AuthenticationServiceImpl) GenerateToken(ctx context.Context, req *authentication.GenerateTokenRequest) (resp *authentication.GenerateTokenResponse, err error) {
	// TODO: Your code here...
	return
}

// CheckToken implements the AuthenticationServiceImpl interface.
func (s *AuthenticationServiceImpl) CheckToken(ctx context.Context, req *authentication.CheckTokenRequest) (resp *authentication.CheckTokenResponse, err error) {
	// TODO: Your code here...
	return
}
