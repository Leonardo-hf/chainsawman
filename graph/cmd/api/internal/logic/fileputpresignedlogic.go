package logic

import (
	"context"

	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilePutPresignedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilePutPresignedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilePutPresignedLogic {
	return &FilePutPresignedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilePutPresignedLogic) FilePutPresigned() (resp *types.PresignedReply, err error) {
	fileId := l.svcCtx.IDGen.New()
	url, err := l.svcCtx.OSSClient.PutPresignedURL(l.ctx, fileId)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("[Graph] fail to generate file presignedurl, err: %v", err)
		return nil, err
	}
	return &types.PresignedReply{
		Url:      url,
		Filename: fileId,
	}, nil
}
