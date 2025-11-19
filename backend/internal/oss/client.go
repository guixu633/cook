package oss

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var OSSClient *oss.Client
var Bucket *oss.Bucket

func InitOSS() {
	var err error
	endpoint := "https://" + os.Getenv("OSS_REGION") + ".aliyuncs.com"
	accessKeyID := os.Getenv("OSS_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("OSS_ACCESS_KEY_SECRET")
	bucketName := os.Getenv("OSS_BUCKET")

	OSSClient, err = oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		log.Fatal("Failed to create OSS client:", err)
	}

	Bucket, err = OSSClient.Bucket(bucketName)
	if err != nil {
		log.Fatal("Failed to get bucket:", err)
	}
	log.Println("OSS client initialized.")
}

func UploadToOSS(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	filename := fmt.Sprintf("meals/%d_%s", time.Now().Unix(), file.Filename)
	err = Bucket.PutObject(filename, src)
	if err != nil {
		return "", err
	}

	// Construct public URL (assuming public read)
	url := fmt.Sprintf("https://%s.%s.aliyuncs.com/%s", os.Getenv("OSS_BUCKET"), os.Getenv("OSS_REGION"), filename)
	return url, nil
}
