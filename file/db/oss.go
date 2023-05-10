package db

import (
	"context"
	"io"
)

type OSSClient interface {
	UploadSource(ctx context.Context, name string, reader io.Reader, size int64) error
	UploadAlgo(ctx context.Context, name string, data []byte) error
	FetchAlgo(ctx context.Context, name string) (io.Reader, error)
}
