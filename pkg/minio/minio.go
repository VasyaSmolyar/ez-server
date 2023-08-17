package minio

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

func Init(cfg *viper.Viper) (*MinioConnect, error) {
	conn := MinioConnect{config: cfg}
	client, err := conn.connect()
	if err != nil {
		return nil, err
	}
	conn.MinioClient = client
	return &conn, nil
}

type MinioConnect struct {
	config      *viper.Viper
	MinioClient *minio.Client
}

func (conn *MinioConnect) connect() (*minio.Client, error) {
	minioClient, err := minio.New(conn.config.GetString("Minio.Endpoint"), &minio.Options{
		Creds: credentials.NewStaticV4(
			conn.config.GetString("Minio.AccessKeyID"),
			conn.config.GetString("Minio.SecretAccessKey"),
			"",
		),
		Secure: conn.config.GetBool("Minio.UseSSL"),
	})
	if err != nil {
		log.Fatal("connect error ", err)
	}

	bucketName := conn.config.GetString("Minio.BucketName")
	location := conn.config.GetString("Minio.Location")

	err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
			return minioClient, nil
		}
		return nil, fmt.Errorf("bucket error %s %s", err, errBucketExists)
	}
	log.Printf("Successfully created %s\n", bucketName)
	return minioClient, nil
}
