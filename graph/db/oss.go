package db

import "context"

type OSSClient interface {
	PutPresignedURL(ctx context.Context, name string) (string, error)
	GetPresignedURL(ctx context.Context, name string) (string, error)
}
