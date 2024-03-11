package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"shorturl/api/internal/config"
	"shorturl/rpc/transform_url/transformurlclient"
)

type ServiceContext struct {
	Config       config.Config
	TransformUrl transformurlclient.TransformUrl
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		TransformUrl: transformurlclient.NewTransformUrl(zrpc.MustNewClient(c.TransformUrl)),
	}
}
