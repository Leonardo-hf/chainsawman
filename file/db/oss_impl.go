package db

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/lifecycle"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
)

type MinioClientImpl struct {
	client *minio.Client
	source string
	algo   string
}

type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	SourceExpired   int
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
	c := &MinioClientImpl{
		client: m,
		source: cfg.SourceBucket,
		algo:   cfg.AlgoBucket,
	}
	err = c.createBucket(cfg.SourceBucket)
	if err != nil {
		logx.Errorf("[OSS] minio create bucket fail, err=%v ", err)
		panic(err)
	}
	err = c.createBucketWithRules(cfg.AlgoBucket, []lifecycle.Rule{
		{
			ID:     "expire-source",
			Status: "Enabled",
			Expiration: lifecycle.Expiration{
				Days: lifecycle.ExpirationDays(cfg.SourceExpired),
			},
		},
	})
	if err != nil {
		logx.Errorf("[OSS] minio create bucket fail, err=%v ", err)
		panic(err)
	}
	logx.Info("[OSS] minio init.")
	return c
}

func (m *MinioClientImpl) createBucket(bucket string) error {
	return m.createBucketWithRules(bucket, []lifecycle.Rule{})
}

func (m *MinioClientImpl) createBucketWithRules(bucket string, rules []lifecycle.Rule) error {
	ctx := context.Background()
	err := m.client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{ObjectLocking: false})
	if err != nil {
		exists, _ := m.client.BucketExists(ctx, bucket)
		if !exists {
			return err
		}
	}
	if len(rules) == 0 {
		return nil
	}
	return m.client.SetBucketLifecycle(ctx, bucket, &lifecycle.Configuration{
		Rules: rules,
	})
}

func (m *MinioClientImpl) upload(ctx context.Context, bucket string, name string, reader io.Reader, size int64) error {
	_, err := m.client.PutObject(ctx, bucket, name, reader, size, minio.PutObjectOptions{})
	return err
}

func (m *MinioClientImpl) fetch(ctx context.Context, bucket string, name string) (io.Reader, error) {
	return m.client.GetObject(ctx, bucket, name, minio.GetObjectOptions{})
}

func (m *MinioClientImpl) UploadSource(ctx context.Context, name string, reader io.Reader, size int64) error {
	return m.upload(ctx, m.source, name, reader, size)
}

func (m *MinioClientImpl) UploadAlgo(ctx context.Context, name string, data []byte) error {
	return m.upload(ctx, m.algo, name, bytes.NewReader(data), int64(len(data)))
}

func (m *MinioClientImpl) FetchAlgo(ctx context.Context, name string) (io.Reader, error) {
	return m.fetch(ctx, m.algo, name)
}
