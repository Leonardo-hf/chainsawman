package logic

import (
	"context"

	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileGetPresignedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileGetPresignedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileGetPresignedLogic {
	return &FileGetPresignedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileGetPresignedLogic) FileGetPresigned(req *types.PresignedRequest) (resp *types.PresignedReply, err error) {
	url, err := l.svcCtx.OSSClient.GetPresignedURL(l.ctx, req.Filename)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("[Graph] fail to generate file presignedurl, err: %v", err)
		return nil, err
	}
	return &types.PresignedReply{
		Url:      url,
		Filename: req.Filename,
	}, nil
}
