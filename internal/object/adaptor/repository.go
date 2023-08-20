package adaptor

import (
	"context"
	"errors"

	"github.com/minio/minio-go/v7"
)

const (
	NotFoundKey = "The specified key does not exist."
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
	if err := objectRepo.minioClient.FGetObject(ctx, *objectRepo.bucketName, objectName, filePath, minio.GetObjectOptions{}); err != nil {
		if err.Error() == NotFoundKey {
			return ErrNotFound
		}
		return err
	}
	return nil
}

func (objectRepo *ObjectRepository) CheckFile(ctx context.Context, objectName string) error {
	if _, err := objectRepo.minioClient.StatObject(ctx, *objectRepo.bucketName, objectName, minio.StatObjectOptions{}); err != nil {
		if err.Error() == NotFoundKey {
			return ErrNotFound
		}
		return err
	}
	// TODO: update TTL
	return nil
}

func (objectRepo *ObjectRepository) DeleteFile(ctx context.Context, objectName string) error {
	if err := objectRepo.minioClient.RemoveObject(ctx, *objectRepo.bucketName, objectName, minio.RemoveObjectOptions{}); err != nil {
		if err.Error() == NotFoundKey {
			return ErrNotFound
		}
		return err
	}
	return nil
}
