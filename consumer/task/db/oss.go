package db

import (
	"context"
	"io"
)

type OSSClient interface {
	Fetch(ctx context.Context, name string) (io.Reader, error)
}
