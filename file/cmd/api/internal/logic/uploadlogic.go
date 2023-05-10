package logic

import (
	"chainsawman/file/cmd/api/internal/svc"
	"chainsawman/file/cmd/api/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

const maxFileSize = 1 << 31

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(r *http.Request) (resp *types.UploadReply, err error) {
	_ = r.ParseMultipartForm(maxFileSize)
	file, header, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	id := l.svcCtx.IDGen.New()
	err = l.svcCtx.OSSClient.UploadSource(l.ctx, id, file, header.Size)
	if err != nil {
		return nil, err
	}
	return &types.UploadReply{
		ID:   id,
		Size: header.Size,
	}, nil
}
