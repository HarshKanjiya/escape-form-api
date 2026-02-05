package services

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/HarshKanjiya/escape-form-api/internal/config"
	"github.com/HarshKanjiya/escape-form-api/internal/storage"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
)

type IUploadService interface {
	GenerateUploadURL(ctx context.Context, userId string, req *types.GenerateUploadURLRequest) (*types.UploadURLResponse, error)
	GenerateDownloadURL(ctx context.Context, userId string, req *types.GenerateDownloadURLRequest) (*types.DownloadURLResponse, error)
	DeleteFile(ctx context.Context, userId string, req *types.DeleteFileRequest) (*types.DeleteFileResponse, error)
}

type UploadService struct {
	cfg *config.Config
}

func NewUploadService(cfg *config.Config) *UploadService {
	return &UploadService{
		cfg: cfg,
	}
}

// GenerateUploadURL creates a presigned URL for file uploads
func (s *UploadService) GenerateUploadURL(ctx context.Context, userId string, req *types.GenerateUploadURLRequest) (*types.UploadURLResponse, error) {

	// Validate file type (add your allowed types)
	allowedTypes := map[string]bool{
		"image/jpeg":       true,
		"image/png":        true,
		"image/gif":        true,
		"image/webp":       true,
		"application/pdf":  true,
		"text/csv":         true,
		"application/json": true,
	}

	if !allowedTypes[req.FileType] {
		return nil, errors.BadRequest("File type not allowed")
	}

	// Set default expiration
	expirationMins := req.ExpirationMins
	if expirationMins == 0 {
		expirationMins = 15 // Default 15 minutes
	}
	if expirationMins > 60 {
		expirationMins = 60 // Max 1 hour
	}

	// Validate intent
	validIntents := map[string]bool{
		"settings": true,
		"response": true,
		"question": true,
		"other":    true,
	}
	if !validIntents[req.Intent] {
		return nil, errors.BadRequest("Invalid intent. Must be one of: settings, response, question, other")
	}

	// Generate unique file key: uploads/form_{formid}/{intent}/{fileName}_{uuid}
	ext := strings.TrimPrefix(filepath.Ext(req.FileName), ".")
	if ext == "" {
		// Extract extension from MIME type
		ext = strings.Split(req.FileType, "/")[1]
	}
	fileNameWithoutExt := strings.TrimSuffix(req.FileName, filepath.Ext(req.FileName))
	fileKey := fmt.Sprintf("uploads/form_%s/%s/%s_%s.%s", req.FormID, req.Intent, fileNameWithoutExt, utils.GenerateUUID(), ext)

	// Generate presigned upload URL
	uploadURL, err := storage.GeneratePresignedUploadURL(ctx, s.cfg.AWS.BucketName, fileKey, expirationMins)
	if err != nil {
		return nil, errors.Internal(err)
	}

	expiresAt := time.Now().Add(time.Duration(expirationMins) * time.Minute).Format(time.RFC3339)

	return &types.UploadURLResponse{
		UploadURL: uploadURL,
		FileKey:   fileKey,
		ExpiresAt: expiresAt,
	}, nil
}

// GenerateDownloadURL creates a presigned URL for file downloads
func (s *UploadService) GenerateDownloadURL(ctx context.Context, userId string, req *types.GenerateDownloadURLRequest) (*types.DownloadURLResponse, error) {

	key := req.FileKey
	if strings.HasPrefix(key, s.cfg.AWS.BucketName+"/") {
		key = strings.TrimPrefix(key, s.cfg.AWS.BucketName+"/")
	}

	// Validate that the file key has the correct format (uploads/form_*/*/)
	if !strings.HasPrefix(key, "uploads/form_") {
		return nil, errors.BadRequest("Invalid file key format")
	}

	// Check if file exists
	exists, err := storage.CheckObjectExists(ctx, s.cfg.AWS.BucketName, key)
	if err != nil {
		return nil, errors.Internal(err)
	}
	if !exists {
		return nil, errors.NotFound("File")
	}

	// Set default expiration
	expirationMins := req.ExpirationMins
	if expirationMins == 0 {
		expirationMins = 60 // Default 1 hour
	}
	if expirationMins > 1440 {
		expirationMins = 1440 // Max 24 hours
	}

	// Generate presigned download URL
	downloadURL, err := storage.GeneratePresignedDownloadURL(ctx, s.cfg.AWS.BucketName, key, expirationMins)
	if err != nil {
		return nil, errors.Internal(err)
	}

	expiresAt := time.Now().Add(time.Duration(expirationMins) * time.Minute).Format(time.RFC3339)

	return &types.DownloadURLResponse{
		DownloadURL: downloadURL,
		FileKey:     req.FileKey,
		ExpiresAt:   expiresAt,
	}, nil
}

// DeleteFile deletes a file from S3
func (s *UploadService) DeleteFile(ctx context.Context, userId string, req *types.DeleteFileRequest) (*types.DeleteFileResponse, error) {

	// Validate that the file key has the correct format (uploads/form_*/*)
	if !strings.HasPrefix(req.FileKey, "uploads/form_") {
		return nil, errors.BadRequest("Invalid file key format")
	}

	// Delete the file
	err := storage.DeleteObject(ctx, s.cfg.AWS.BucketName, req.FileKey)
	if err != nil {
		return nil, errors.Internal(err)
	}

	return &types.DeleteFileResponse{
		Success: true,
		Message: "File deleted successfully",
	}, nil
}
