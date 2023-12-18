package db

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/lifecycle"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/url"
	"time"
)

type MinioClientImpl struct {
	client  *minio.Client
	source  string
	algo    string
	lib     string
	expired time.Duration
}

type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	SourceExpired   int
	URLExpired      int
	SourceBucket    string
	LibBucket       string
	AlgoBucket      string
}

func InitMinioClient(cfg *MinioConfig) OSSClient {
	m, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Secure: cfg.UseSSL})
	if err != nil {
		logx.Errorf("[OSS] minio connect fail, err=%v ", err)
		panic(err)
	}
	source := cfg.SourceBucket
	c := &MinioClientImpl{
		client:  m,
		source:  source,
		lib:     cfg.LibBucket,
		algo:    cfg.AlgoBucket,
		expired: time.Duration(cfg.URLExpired) * time.Second,
	}
	ctx := context.Background()
	if ok, _ := m.BucketExists(ctx, source); !ok {
		err = m.MakeBucket(ctx, source, minio.MakeBucketOptions{ObjectLocking: false})
		if err != nil {
			logx.Errorf("[Graph] fail to create bucket in minio, err: %v", err)
			panic(err)
		}
		err = m.SetBucketLifecycle(ctx, source, &lifecycle.Configuration{
			Rules: []lifecycle.Rule{
				{
					ID:     "expire-source",
					Status: "Enabled",
					Expiration: lifecycle.Expiration{
						Days: lifecycle.ExpirationDays(cfg.SourceExpired),
					},
				},
			},
		})
		if err != nil {
			logx.Errorf("[Graph] fail to create bucket in minio, err: %v", err)
			panic(err)
		}
	}
	return c
}

func (m *MinioClientImpl) PutSourcePresignedURL(ctx context.Context, name string) (string, error) {
	presignedURL, err := m.client.PresignedPutObject(ctx, m.source, name, m.expired)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

func (m *MinioClientImpl) PutLibPresignedURL(ctx context.Context, name string) (string, error) {
	presignedURL, err := m.client.PresignedPutObject(ctx, m.lib, name, m.expired)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

func (m *MinioClientImpl) GetAlgoPresignedURL(ctx context.Context, name string) (string, error) {
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%v\"", name))
	presignedURL, err := m.client.PresignedGetObject(ctx, m.algo, name, m.expired, reqParams)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

func (m *MinioClientImpl) FetchAlgo(ctx context.Context, name string) (io.Reader, error) {
	return m.client.GetObject(ctx, m.algo, name, minio.GetObjectOptions{})
}
