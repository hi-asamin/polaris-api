package infrastructure

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var s3Client *s3.Client
var s3Once sync.Once

// NewS3Client はS3クライアントを初期化します
func NewS3Client() {
	s3Once.Do(func() {
		log.Print("Start create new s3 client")
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatalf("failed to load AWS config: %v", err)
		}

		s3Client = s3.NewFromConfig(cfg)
		log.Print("Complete create new s3 client")
	})
}

// GetS3Client は初期化されたS3クライアントを返します
func GetS3Client() *s3.Client {
	if s3Client != nil {
		NewS3Client()
	}
	return s3Client
}

func GetS3BucketName() string {
	return os.Getenv("S3_BUCKET_NAME")
}
