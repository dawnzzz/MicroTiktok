// Code generated by Kitex v0.5.2. DO NOT EDIT.

package idgeneratorservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	id_generator "github.com/dawnzzz/MicroTiktok/kitex_gen/id_generator"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetUserID(ctx context.Context, req *id_generator.GetUserIdRequest, callOptions ...callopt.Option) (r *id_generator.GetUserIdResponse, err error)
	GetVideoID(ctx context.Context, req *id_generator.GetVideoIdRequest, callOptions ...callopt.Option) (r *id_generator.GetVideoIdResponse, err error)
	GetCommentID(ctx context.Context, req *id_generator.GetCommentIdRequest, callOptions ...callopt.Option) (r *id_generator.GetCommentIdResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kIdGeneratorServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kIdGeneratorServiceClient struct {
	*kClient
}

func (p *kIdGeneratorServiceClient) GetUserID(ctx context.Context, req *id_generator.GetUserIdRequest, callOptions ...callopt.Option) (r *id_generator.GetUserIdResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserID(ctx, req)
}

func (p *kIdGeneratorServiceClient) GetVideoID(ctx context.Context, req *id_generator.GetVideoIdRequest, callOptions ...callopt.Option) (r *id_generator.GetVideoIdResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVideoID(ctx, req)
}

func (p *kIdGeneratorServiceClient) GetCommentID(ctx context.Context, req *id_generator.GetCommentIdRequest, callOptions ...callopt.Option) (r *id_generator.GetCommentIdResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCommentID(ctx, req)
}
