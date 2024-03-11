package logic

import (
	"context"

	"shorturl/rpc/transform_url/internal/svc"
	"shorturl/rpc/transform_url/transform_url"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *transform_url.ExpandReq) (*transform_url.ExpandResp, error) {
	longUrl, err := l.svcCtx.Model.FindOne(l.ctx, in.Shorten)
	if err != nil {
		return nil, err
	}
	return &transform_url.ExpandResp{
		Url: longUrl.Url,
	}, nil
}
