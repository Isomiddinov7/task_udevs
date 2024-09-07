package utils

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

// imageMimeTypes maps file extensions to their MIME types
var imageMimeTypes = map[string]string{
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".png":  "image/png",
	".gif":  "image/gif",
	".bmp":  "image/bmp",
	".webp": "image/webp",
}

// UploadImage uploads a file to an AWS S3 bucket with metadata.
func UploadImage(file *multipart.FileHeader) (string, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	awsRegion := os.Getenv("AWS_REGION")
	if awsRegion == "" {
		log.Fatalf("AWS_REGION is not set in the environment")
	}

	bucketName := os.Getenv("AWS_S3_BUCKET")
	if bucketName == "" {
		log.Fatalf("AWS_S3_BUCKET is not set in the environment")
	}

	// Load the AWS config with the specified region
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsRegion))
	if err != nil {
		log.Printf("Error loading AWS config: %v", err)
		return "", fmt.Errorf("Error loading AWS config: %w", err)
	}

	// Create a new S3 client
	client := s3.NewFromConfig(cfg)

	// Ensure the bucket is in the correct region by checking its location
	location, err := client.GetBucketLocation(context.TODO(), &s3.GetBucketLocationInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Printf("Error getting bucket location: %v", err)
		return "", fmt.Errorf("Error getting bucket location: %w", err)
	}

	// Determine the bucket's region
	bucketRegion := string(location.LocationConstraint)
	if bucketRegion == "" {
		bucketRegion = "us-east-1" // Default for the "US East (N. Virginia)" region
	}

	// Reload the AWS config if the bucket region is different
	if bucketRegion != awsRegion {
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithRegion(bucketRegion))
		if err != nil {
			log.Printf("Error reloading AWS config with bucket region: %v", err)
			return "", fmt.Errorf("Error reloading AWS config with bucket region: %w", err)
		}
		client = s3.NewFromConfig(cfg)
	}

	// Open the file
	f, err := file.Open()
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return "", fmt.Errorf("Error opening file: %w", err)
	}
	defer f.Close()

	// Determine Content-Type based on file extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	contentType, ok := imageMimeTypes[ext]
	if !ok {
		contentType = "application/octet-stream" // Default content type if not found
	}

	// Create an uploader with the S3 client
	uploader := manager.NewUploader(client)

	key := file.Filename
	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(key),
		Body:        f,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		log.Printf("Error uploading file to S3: %v", err)
		return "", fmt.Errorf("Error uploading file to S3: %w", err)
	}

	// Return the URL of the uploaded file
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, bucketRegion, key)
	fmt.Println(url)
	return url, nil
}
