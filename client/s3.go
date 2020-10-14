package client

import (
	"io"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"

	"github.com/YusukeKishino/go-blog/config"
)

type S3Client interface {
	UploadImage(data io.Reader, filename string) (string, error)
	Upload(data io.Reader, key string) (string, error)
}

type s3Client struct {
	conf *config.AppConfig
}

func NewS3Client(conf *config.AppConfig) S3Client {
	return &s3Client{conf: conf}
}

func (c *s3Client) UploadImage(data io.Reader, filename string) (string, error) {
	return c.Upload(data, filepath.Join("images/", filename))
}

func (c *s3Client) Upload(data io.Reader, key string) (string, error) {
	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess)

	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(c.conf.S3Bucket),
		Key:    aws.String(key),
		Body:   data,
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to upload file")
	}
	return output.Location, nil
}
