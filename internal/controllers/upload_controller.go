package controllers

import (
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UploadController struct {
	validator     *validator.Validate
	uploadService services.IUploadService
}

func NewUploadController(service services.IUploadService) *UploadController {
	return &UploadController{
		validator:     validator.New(),
		uploadService: service,
	}
}

// @Summary Generate presigned upload URL
// @Description Generate a presigned URL for uploading a file to S3
// @Tags upload
// @Accept json
// @Produce json
// @Param request body types.GenerateUploadURLRequest true "Upload URL request"
// @Success 200 {object} types.UploadURLResponse
// @Router /upload/generate-url [post]
func (uc *UploadController) GenerateUploadURL(c *fiber.Ctx) error {
	userId, ok := utils.GetUserId(c)
	if !ok {
		return errors.Unauthorized("")
	}

	var req types.GenerateUploadURLRequest
	if err := c.BodyParser(&req); err != nil {
		return errors.BadRequest("Invalid request body")
	}

	if err := uc.validator.Struct(&req); err != nil {
		return errors.BadRequest(err.Error())
	}

	response, err := uc.uploadService.GenerateUploadURL(c.Context(), userId, &req)
	if err != nil {
		return err
	}

	return utils.Success(c, response, "Upload URL generated successfully")
}

// @Summary Generate presigned download URL
// @Description Generate a presigned URL for downloading a file from S3
// @Tags upload
// @Accept json
// @Produce json
// @Param request body types.GenerateDownloadURLRequest true "Download URL request"
// @Success 200 {object} types.DownloadURLResponse
// @Router /upload/download-url [post]
func (uc *UploadController) GenerateDownloadURL(c *fiber.Ctx) error {
	userId, ok := utils.GetUserId(c)
	if !ok {
		return errors.Unauthorized("")
	}

	var req types.GenerateDownloadURLRequest
	if err := c.BodyParser(&req); err != nil {
		return errors.BadRequest("Invalid request body")
	}

	if err := uc.validator.Struct(&req); err != nil {
		return errors.BadRequest(err.Error())
	}

	response, err := uc.uploadService.GenerateDownloadURL(c.Context(), userId, &req)
	if err != nil {
		return err
	}

	return utils.Success(c, response, "Download URL generated successfully")
}

// @Summary Delete file
// @Description Delete a file from S3
// @Tags upload
// @Accept json
// @Produce json
// @Param request body types.DeleteFileRequest true "Delete file request"
// @Success 200 {object} types.DeleteFileResponse
// @Router /upload/delete [delete]
func (uc *UploadController) DeleteFile(c *fiber.Ctx) error {
	userId, ok := utils.GetUserId(c)
	if !ok {
		return errors.Unauthorized("")
	}

	var req types.DeleteFileRequest
	if err := c.BodyParser(&req); err != nil {
		return errors.BadRequest("Invalid request body")
	}

	if err := uc.validator.Struct(&req); err != nil {
		return errors.BadRequest(err.Error())
	}

	response, err := uc.uploadService.DeleteFile(c.Context(), userId, &req)
	if err != nil {
		return err
	}

	return utils.Success(c, response, "File deleted successfully")
}
