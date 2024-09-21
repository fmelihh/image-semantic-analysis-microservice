package db

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewFileStorage() (*minio.Client, error) {
	minioClient, err := minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("CD3X6NS331R7PMWQ51RB", "AZm1+E+pGwWOf4fgRuX9tvSBkk+28uX+ceAKhLPu", ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalf("Error initializing Minio client: %v", err)
		return nil, err
	}

	err = createBucket("image", minioClient)
	if err != nil {
		log.Fatalf("Error creating bucket: %v", err)
		return nil, err
	}
	return minioClient, nil
}

func createBucket(bucketName string, minioClient *minio.Client) error {
	ctx := context.Background()
	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "us-east-1"})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("Bucket already exists")
		} else {
			return err
		}
	}
	log.Printf("Successfully created %s\n", bucketName)
	return nil
}
