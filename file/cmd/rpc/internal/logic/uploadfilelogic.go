package logic

import (
	"chainsawman/file/cmd/rpc/internal/svc"
	"chainsawman/file/cmd/rpc/types/rpc/file"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadFileLogic) UploadFile(in *file.FileUploadReq) (*file.FileIDReply, error) {
	id := in.Name + l.svcCtx.IDGen.New() + ".csv"
	err := l.svcCtx.OSSClient.UploadAlgo(l.ctx, id, in.Data)
	if err != nil {
		return nil, err
	}
	return &file.FileIDReply{
		Id: id,
	}, nil
}
