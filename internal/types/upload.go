package types

// Request structs
type GenerateUploadURLRequest struct {
	FileName       string `json:"fileName" validate:"required"`
	FileType       string `json:"fileType" validate:"required"`
	ExpirationMins int64  `json:"expirationMins,omitempty"`
}

type GenerateDownloadURLRequest struct {
	FileKey        string `json:"fileKey" validate:"required"`
	ExpirationMins int64  `json:"expirationMins,omitempty"`
}

type DeleteFileRequest struct {
	FileKey string `json:"fileKey" validate:"required"`
}

// Response structs
type UploadURLResponse struct {
	UploadURL string `json:"uploadUrl"`
	FileKey   string `json:"fileKey"`
	ExpiresAt string `json:"expiresAt"`
}

type DownloadURLResponse struct {
	DownloadURL string `json:"downloadUrl"`
	FileKey     string `json:"fileKey"`
	ExpiresAt   string `json:"expiresAt"`
}

type DeleteFileResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
