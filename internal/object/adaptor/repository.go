package adaptor

import (
	"context"
	"errors"

	"github.com/minio/minio-go/v7"
)

var ErrNotFound error = errors.New("resource was not found")

func Init(bucketName string, minioClient *minio.Client) *ObjectRepository {
	return &ObjectRepository{bucketName: &bucketName, minioClient: minioClient}
}

type ObjectRepository struct {
	bucketName  *string
	minioClient *minio.Client
}

func (objectRepo *ObjectRepository) UploadFile(ctx context.Context, objectName, filePath, contentType string) error {
	_, err := objectRepo.minioClient.FPutObject(ctx, *objectRepo.bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	return err
}

func (objectRepo *ObjectRepository) DownloadFile(ctx context.Context, objectName, filePath string) error {
	return objectRepo.minioClient.FGetObject(ctx, *objectRepo.bucketName, objectName, filePath, minio.GetObjectOptions{})
}
