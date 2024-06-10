package db

import (
	"context"
	"io"
)

type OSSClient interface {
	FetchSource(ctx context.Context, name string) (io.Reader, error)

	FetchAlgo(ctx context.Context, name string) (io.Reader, error)

	AddSource(ctx context.Context, name string, content []byte) (string, error)
}
