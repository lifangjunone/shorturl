// Code generated by goctl. DO NOT EDIT.
// Source: transform_url.proto

package transformurlclient

import (
	"context"

	"shorturl/rpc/transform_url/transform_url"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ExpandReq   = transform_url.ExpandReq
	ExpandResp  = transform_url.ExpandResp
	ShortenReq  = transform_url.ShortenReq
	ShortenResp = transform_url.ShortenResp

	TransformUrl interface {
		Shorten(ctx context.Context, in *ShortenReq, opts ...grpc.CallOption) (*ShortenResp, error)
		Expand(ctx context.Context, in *ExpandReq, opts ...grpc.CallOption) (*ExpandResp, error)
	}

	defaultTransformUrl struct {
		cli zrpc.Client
	}
)

func NewTransformUrl(cli zrpc.Client) TransformUrl {
	return &defaultTransformUrl{
		cli: cli,
	}
}

func (m *defaultTransformUrl) Shorten(ctx context.Context, in *ShortenReq, opts ...grpc.CallOption) (*ShortenResp, error) {
	client := transform_url.NewTransformUrlClient(m.cli.Conn())
	return client.Shorten(ctx, in, opts...)
}

func (m *defaultTransformUrl) Expand(ctx context.Context, in *ExpandReq, opts ...grpc.CallOption) (*ExpandResp, error) {
	client := transform_url.NewTransformUrlClient(m.cli.Conn())
	return client.Expand(ctx, in, opts...)
}
