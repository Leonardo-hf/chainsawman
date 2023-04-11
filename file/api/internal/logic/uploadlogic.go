package logic

import (
	"context"
	"io"
	"net/http"
	"os"
	p "path"
	"time"

	"chainsawman/file/api/internal/svc"
	"chainsawman/file/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	c := l.svcCtx.Config
	_ = r.ParseMultipartForm(maxFileSize)
	file, header, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	id := l.svcCtx.IDGenerator.String()
	path := p.Join(c.Path, id+".csv")
	target, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	_, err = io.Copy(target, file)
	if err != nil {
		return nil, err
	}
	// 延迟删除文件
	time.AfterFunc(time.Duration(c.Expired)*time.Second, func() {
		if err = os.Remove(path); err != nil {
			logx.Errorf("[file] remove fail, err: %v, id: %v", err, id)
		}
	})
	return &types.UploadReply{
		ID:   id,
		Size: header.Size,
	}, nil
}
