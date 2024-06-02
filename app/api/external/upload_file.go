package external

import (
	"context"
	"fmt"
	"mime/multipart"
	"museum/app/utils"
	"museum/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	contentType    = "application/octet-stream"
	lengthFileName = 6
)

type UploadFileAPI struct {
	bucketName string
	endpoint   string

	file     *multipart.FileHeader
	s3Client *minio.Client
}

func NewUploadFileAPI(file *multipart.FileHeader) *UploadFileAPI {
	return &UploadFileAPI{
		file: file,
	}
}

func (c *UploadFileAPI) UploadObject() (string, error) {
	if err := c.initClient(); err != nil {
		return "", err
	}

	object, err := c.file.Open()
	if err != nil {
		return "", err
	}
	defer object.Close()

	fileName := c.generateFileName()
	info, err := c.s3Client.PutObject(
		context.Background(), c.bucketName, fileName,
		object, c.file.Size, minio.PutObjectOptions{ContentType: contentType},
	)

	if err != nil {
		return "", err
	}

	return c.objectURL(&info), nil
}

func (c *UploadFileAPI) objectURL(info *minio.UploadInfo) string {
	return fmt.Sprintf(
		"https://%s/%s/%s", c.endpoint, info.Bucket, info.Key,
	)
}

func (c *UploadFileAPI) initClient() error {
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

func (c *UploadFileAPI) getConf() (config.StorageS3, error) {
	cfg, err := config.GetConf()

	return cfg.StorageS3, err
}

func (c *UploadFileAPI) generateFileName() string {
	return utils.GeneratePassword(lengthFileName, true, false, true)
}
