package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/hash"
	"shorturl/rpc/transform_url/model"

	"shorturl/rpc/transform_url/internal/svc"
	"shorturl/rpc/transform_url/transform_url"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *transform_url.ShortenReq) (*transform_url.ShortenResp, error) {
	shortUrl := hash.Md5Hex([]byte(in.Url))[:6]
	_, err := l.svcCtx.Model.Insert(l.ctx, &model.Shorturl{
		Shorten: shortUrl,
		Url:     in.Url,
	})
	if err != nil {
		return nil, err
	}
	return &transform_url.ShortenResp{
		Shorten: shortUrl,
	}, nil
}
