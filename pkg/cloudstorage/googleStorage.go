package cloudstorage

import (
	"bytes"
	"context"
	"os"

	"cloud.google.com/go/storage"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

// GoogleStorageBucket helper class for Google Cloud Buckets
type GoogleStorageBucket struct {
	client     *storage.Client
	bucketName string

	logger *zap.SugaredLogger
}

// NewGoogleStorageBucket initializes a GoogleStorageBucket
func NewGoogleStorageBucket(bucketName string) (*GoogleStorageBucket, error) {
	credPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	client, err := storage.NewClient(context.TODO(), option.WithCredentialsFile(credPath))

	if err != nil {
		return nil, err
	}

	return &GoogleStorageBucket{client: client, bucketName: bucketName}, nil
}

// Close - closes the Client. Close need not be called at program exit.
func (g *GoogleStorageBucket) Close() error {
	return g.client.Close()
}

func (g *GoogleStorageBucket) UploadFile(fileName string, bytes bytes.Buffer) error {
	/*
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
		defer cancel()

		objectWriter := g.client.Bucket(g.bucketName).Object(fileName).NewWriter(ctx)
	*/

	return nil
}

func (g *GoogleStorageBucket) DownloadFile(fileName string) (bytes.Buffer, error) {
	return bytes.Buffer{}, nil
}
