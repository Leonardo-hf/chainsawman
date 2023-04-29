package logic

import (
	"context"
	"os"
	p "path"
	"time"

	"chainsawman/file/cmd/rpc/internal/svc"
	"chainsawman/file/cmd/rpc/types/rpc"

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

func (l *UploadFileLogic) UploadFile(in *rpc.FileUploadReq) (*rpc.FileIDReply, error) {
	c := l.svcCtx.Config
	id := l.svcCtx.IDGen.New()
	name := in.Name + id
	path := p.Join(c.Path, name+".csv")
	target, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	_, err = target.Write(in.Data)
	if err != nil {
		return nil, err
	}
	if in.Expired > 0 {
		// 延迟删除文件
		time.AfterFunc(time.Duration(in.Expired)*time.Second, func() {
			if err = os.Remove(path); err != nil {
				logx.Errorf("[file] remove fail, err: %v, id: %v", err, name)
			}
		})
	}
	return &rpc.FileIDReply{
		Id: name,
	}, nil
}
