package db

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/lifecycle"
	"github.com/zeromicro/go-zero/core/logx"
	"net/url"
	"time"
)

type MinioClientImpl struct {
	client  *minio.Client
	source  string
	algo    string
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
	bucket := cfg.SourceBucket
	c := &MinioClientImpl{
		client:  m,
		source:  bucket,
		algo:    cfg.AlgoBucket,
		expired: time.Duration(cfg.URLExpired) * time.Second,
	}
	ctx := context.Background()
	if ok, _ := m.BucketExists(ctx, bucket); !ok {
		err = m.MakeBucket(ctx, bucket, minio.MakeBucketOptions{ObjectLocking: false})
		if err != nil {
			logx.Errorf("[Graph] fail to create bucket in minio, err: %v", err)
			panic(err)
		}
		err = m.SetBucketLifecycle(ctx, bucket, &lifecycle.Configuration{
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

func (m *MinioClientImpl) PutPresignedURL(ctx context.Context, name string) (string, error) {
	presignedURL, err := m.client.PresignedPutObject(ctx, m.source, name, m.expired)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

func (m *MinioClientImpl) GetPresignedURL(ctx context.Context, name string) (string, error) {
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%v\"", name))
	presignedURL, err := m.client.PresignedGetObject(ctx, m.algo, name, m.expired, reqParams)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}
