package db

import "context"

type OSSClient interface {
	PutSourcePresignedURL(ctx context.Context, name string) (string, error)
	PutLibPresignedURL(ctx context.Context, name string) (string, error)
	GetAlgoPresignedURL(ctx context.Context, name string) (string, error)
}
