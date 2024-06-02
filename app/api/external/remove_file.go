package external

import (
	"context"
	"fmt"
	"museum/config"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type RemoveFileAPI struct {
	bucketName string
	endpoint   string

	fileURL  string
	s3Client *minio.Client
}

func NewRemoveFileAPI(fileURL string) *RemoveFileAPI {
	return &RemoveFileAPI{
		fileURL: fileURL,
	}
}

func (c *RemoveFileAPI) RemoveObject() error {
	if err := c.initClient(); err != nil {
		return err
	}

	objectName, err := c.nameObject()
	if err != nil {
		return err
	}

	err = c.s3Client.RemoveBucket(
		context.Background(), objectName,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *RemoveFileAPI) initClient() error {
	var cfg config.StorageS3
	var err error

	cfg, err = c.getConf()
	if err != nil {
		return err
	}

	c.bucketName = cfg.BucketName
	c.endpoint = cfg.EndPoint

	c.s3Client, err = minio.New(
		c.endpoint,
		&minio.Options{
			Creds: credentials.NewStaticV4(
				cfg.AccessKey, cfg.SecretAccessKey, "",
			),
			Region: cfg.Region,
			Secure: true,
		},
	)

	return err
}

func (c *RemoveFileAPI) getConf() (config.StorageS3, error) {
	cfg, err := config.GetConf()

	return cfg.StorageS3, err
}

func (c *RemoveFileAPI) nameObject() (string, error) {
	separatedStr := strings.Split(c.fileURL, "/")
	if len(separatedStr) == 0 {
		return "", fmt.Errorf("Ошибка парсинга url строки: %s", c.fileURL)
	}

	return separatedStr[len(separatedStr)-1], nil
}
