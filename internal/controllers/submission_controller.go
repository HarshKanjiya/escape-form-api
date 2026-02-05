package controllers

import (
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type SubmissionController struct {
	submissionService services.ISubmissionService
}

func NewSubmissionController(service services.ISubmissionService) *SubmissionController {
	return &SubmissionController{
		submissionService: service,
	}
}

// @Summary Get form by domain
// @Description Retrieve form details by domain
// @Tags submission
// @Accept json
// @Produce json
// @Param domain path string true "Form Domain"
// @Success 200 {object} int64
// @Router /submission/{domain} [get]
func (pc *SubmissionController) GetForm(c *fiber.Ctx) error {

	domain := c.Params("domain", "")
	if domain == "" {
		return errors.BadRequest("Id is required")
	}

	form, err := pc.submissionService.GetForm(c.Context(), domain)
	if err != nil {
		return err
	}

	return utils.Success(c, form, "Form fetched successfully")
}
