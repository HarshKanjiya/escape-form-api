package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/HarshKanjiya/escape-form-api/internal/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Global S3Client
var S3Client *s3.Client

var S3PresignClient *s3.PresignClient

func Connect(cfg *config.Config) error {
	if cfg.AWS.AccessKey == "" || cfg.AWS.SecretKey == "" {
		return fmt.Errorf("AWS credentials not configured")
	}

	awsCfg, err := awsconfig.LoadDefaultConfig(context.Background(),
		awsconfig.WithRegion(cfg.AWS.Region),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AWS.AccessKey,
			cfg.AWS.SecretKey,
			"",
		)),
	)
	if err != nil {
		return fmt.Errorf("failed to load AWS config: %w", err)
	}

	s3Options := func(o *s3.Options) {
		if cfg.AWS.EndPoint != "" {
			o.BaseEndpoint = aws.String(cfg.AWS.EndPoint)
			o.UsePathStyle = true // Required for custom endpoints (Cloudflare R2)
		}
	}

	// Initialize S3 client
	S3Client = s3.NewFromConfig(awsCfg, s3Options)

	// Initialize Presign client
	S3PresignClient = s3.NewPresignClient(S3Client)

	log.Println("AWS S3 client connected successfully")
	return nil
}

func GeneratePresignedUploadURL(ctx context.Context, bucketName, key string, expirationMinutes int64) (string, error) {
	if S3PresignClient == nil {
		return "", fmt.Errorf("S3 presign client not initialized")
	}

	request, err := S3PresignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(expirationMinutes) * time.Minute
	})

	if err != nil {
		return "", fmt.Errorf("failed to generate presigned upload URL: %w", err)
	}

	return request.URL, nil
}

func GeneratePresignedDownloadURL(ctx context.Context, bucketName, key string, expirationMinutes int64) (string, error) {
	if S3PresignClient == nil {
		return "", fmt.Errorf("S3 presign client not initialized")
	}

	request, err := S3PresignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(expirationMinutes) * time.Minute
	})

	if err != nil {
		return "", fmt.Errorf("failed to generate presigned download URL: %w", err)
	}

	return request.URL, nil
}

// DeleteObject deletes an object from S3
func DeleteObject(ctx context.Context, bucketName, key string) error {
	if S3Client == nil {
		return fmt.Errorf("S3 client not initialized")
	}

	_, err := S3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})

	if err != nil {
		return fmt.Errorf("failed to delete object: %w", err)
	}

	return nil
}

func CheckObjectExists(ctx context.Context, bucketName, key string) (bool, error) {
	if S3Client == nil {
		return false, fmt.Errorf("S3 client not initialized")
	}

	_, err := S3Client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})

	if err != nil {
		return false, nil
	}

	return true, nil
}
