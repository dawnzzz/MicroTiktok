// Code generated by Kitex v0.5.2. DO NOT EDIT.

package authenticationservice

import (
	server "github.com/cloudwego/kitex/server"
	authentication "github.com/dawnzzz/MicroTiktok/kitex_gen/authentication"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler authentication.AuthenticationService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
