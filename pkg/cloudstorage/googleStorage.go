package cloudstorage

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

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
func NewGoogleStorageBucket(bucketName string, logger *zap.SugaredLogger) (*GoogleStorageBucket, error) {
	credPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	client, err := storage.NewClient(context.TODO(), option.WithCredentialsFile(credPath))

	if err != nil {
		return nil, err
	}

	return &GoogleStorageBucket{client: client, bucketName: bucketName, logger: logger}, nil
}

// Close - closes the Client. Close need not be called at program exit.
func (g *GoogleStorageBucket) Close() error {
	return g.client.Close()
}

func (g *GoogleStorageBucket) UploadFile(fileName string, bytes bytes.Buffer) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()

	objectWriter := g.client.Bucket(g.bucketName).Object(fileName).NewWriter(ctx)

	defer objectWriter.Close()

	n, err := objectWriter.Write(bytes.Bytes())

	if err != nil {
		g.logger.Debugf("Error writing to gcloud bucket: %v", err)
		return err
	}

	g.logger.Infof("Uploaded %v bytes in %v file to gcloud", n, fileName)

	return nil
}

func (g *GoogleStorageBucket) DownloadFile(fileName string) (bytes.Buffer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	rc, err := g.client.Bucket(g.bucketName).Object(fileName).NewReader(ctx)
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("Object(%q).NewReader: %v", fileName, err)
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("ioutil.ReadAll: %v", err)
	}
	return *bytes.NewBuffer(data), nil
}
