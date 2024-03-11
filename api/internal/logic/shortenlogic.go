package logic

import (
	"context"
	"shorturl/rpc/transform_url/transform_url"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req *types.ShortenReq) (*types.ShortenResp, error) {
	resp, err := l.svcCtx.TransformUrl.Shorten(l.ctx, &transform_url.ShortenReq{
		Url: req.Url,
	})
	if err != nil {
		return nil, err
	}
	return &types.ShortenResp{
		Shorten: resp.Shorten,
	}, nil
}
