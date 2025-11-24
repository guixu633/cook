package oss

import (
	"fmt"
	"io"
	"path"
	"server/config"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Client struct {
	bucket   *oss.Bucket
	basePath string
	endpoint string
}

func NewClient(cfg *config.Config) (*Client, error) {
	client, err := oss.New(cfg.OSS.Endpoint, cfg.OSS.AccessKeyID, cfg.OSS.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(cfg.OSS.BucketName)
	if err != nil {
		return nil, err
	}

	return &Client{
		bucket:   bucket,
		basePath: cfg.OSS.BasePath,
		endpoint: cfg.OSS.Endpoint,
	}, nil
}

// UploadToTmp uploads a file to the temporary directory in OSS
// Returns the object key (path in OSS) and error
func (c *Client) UploadToTmp(filename string, reader io.Reader) (string, error) {
	objectKey := path.Join(c.basePath, "tmp", filename)
	err := c.bucket.PutObject(objectKey, reader)
	if err != nil {
		return "", err
	}
	return filename, nil // We return just the filename as requested by the user flow
}

// MoveFromTmpToMeal moves a file from tmp directory to meal specific directory
// Returns the public URL of the new file
func (c *Client) MoveFromTmpToMeal(filename string, mealID uint) (string, error) {
	srcObjectKey := path.Join(c.basePath, "tmp", filename)
	destObjectKey := path.Join(c.basePath, fmt.Sprintf("%d", mealID), filename)

	// Copy object
	_, err := c.bucket.CopyObject(srcObjectKey, destObjectKey)
	if err != nil {
		return "", fmt.Errorf("failed to copy object: %w", err)
	}

	// Delete source object
	err = c.bucket.DeleteObject(srcObjectKey)
	if err != nil {
		// Log error but don't fail the operation since copy was successful
		fmt.Printf("Warning: failed to delete temp file %s: %v\n", srcObjectKey, err)
	}

	// Generate public URL
	// Assuming public read bucket or standard OSS URL format
	// Format: http://<bucket>.<endpoint>/<objectKey>
	// We need to strip http:// from endpoint if present for clean formatting if we construct manually,
	// or use the endpoint as is.

	// Clean endpoint: remove http:// or https://
	endpoint := c.endpoint
	if strings.HasPrefix(endpoint, "http://") {
		endpoint = strings.TrimPrefix(endpoint, "http://")
	} else if strings.HasPrefix(endpoint, "https://") {
		endpoint = strings.TrimPrefix(endpoint, "https://")
	}

	url := fmt.Sprintf("http://%s.%s/%s", c.bucket.BucketName, endpoint, destObjectKey)
	return url, nil
}
