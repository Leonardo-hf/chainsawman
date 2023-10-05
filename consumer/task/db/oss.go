package db

import (
	"context"
	"io"
)

type OSSClient interface {
	FetchSource(ctx context.Context, name string) (io.Reader, error)

	FetchAlgo(ctx context.Context, name string) (io.Reader, error)

	AlgoGenerated(ctx context.Context, name string) bool
}
