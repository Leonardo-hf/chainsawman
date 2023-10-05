package logic

import (
	"context"
	"github.com/google/uuid"

	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileLibPutPresignedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileLibPutPresignedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileLibPutPresignedLogic {
	return &FileLibPutPresignedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileLibPutPresignedLogic) FileLibPutPresigned() (resp *types.PresignedReply, err error) {
	fileId := uuid.New().String()
	url, err := l.svcCtx.OSSClient.PutLibPresignedURL(l.ctx, fileId)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("[Graph] fail to generate file presignedurl, err: %v", err)
		return nil, err
	}
	return &types.PresignedReply{
		Url:      url,
		Filename: fileId,
	}, nil
}