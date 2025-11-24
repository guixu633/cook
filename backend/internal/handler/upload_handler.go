package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"server/internal/pkg/oss"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UploadHandler struct {
	ossClient *oss.Client
}

func NewUploadHandler(ossClient *oss.Client) *UploadHandler {
	return &UploadHandler{ossClient: ossClient}
}

func (h *UploadHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer src.Close()

	// Generate filename: timestamp + random 8 char uuid
	ext := filepath.Ext(file.Filename)
	randomStr := uuid.New().String()[:8]
	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("%d_%s%s", timestamp, randomStr, ext)

	// Upload to OSS tmp
	_, err = h.ossClient.UploadToTmp(filename, src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload to storage"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"filename": filename,
		"url":      filename, // Frontend expects filename to send back later
	})
}
