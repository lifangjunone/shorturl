package logic

import (
	"context"
	"shorturl/rpc/transform_url/transform_url"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandLogic) Expand(req *types.ExpandReq) (*types.ExpandResp, error) {
	// todo: add your logic here and delete this line
	longUrl, err := l.svcCtx.TransformUrl.Expand(l.ctx, &transform_url.ExpandReq{Shorten: req.Shorten})
	if err != nil {
		return nil, err
	}
	return &types.ExpandResp{
		Url: longUrl.Url,
	}, nil
}
