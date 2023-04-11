// Code generated by goctl. DO NOT EDIT.
// Source: file.proto

package file

import (
	"context"

	"chainsawman/file/rpc/types/rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FileReply = rpc.FileReply
	IDReq     = rpc.IDReq

	File interface {
		FetchFile(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*FileReply, error)
	}

	defaultFile struct {
		cli zrpc.Client
	}
)

func NewFile(cli zrpc.Client) File {
	return &defaultFile{
		cli: cli,
	}
}

func (m *defaultFile) FetchFile(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*FileReply, error) {
	client := rpc.NewFileClient(m.cli.Conn())
	return client.FetchFile(ctx, in, opts...)
}
