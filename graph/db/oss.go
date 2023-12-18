package db

import (
	"context"
	"io"
)

type OSSClient interface {
	FetchAlgo(ctx context.Context, name string) (io.Reader, error)
	PutSourcePresignedURL(ctx context.Context, name string) (string, error)
	PutLibPresignedURL(ctx context.Context, name string) (string, error)
	GetAlgoPresignedURL(ctx context.Context, name string) (string, error)
}
