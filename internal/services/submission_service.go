package services

import (
	"context"
	"encoding/json"

	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
)

type ISubmissionService interface {
	GetForm(ctx context.Context, domain string) (*types.GetSubmissionFormResponse, error)
}

type SubmissionService struct {
	formRepo        repositories.IFormRepo
	formVersionRepo repositories.IFormVersionRepo
}

func NewSubmissionService(formRepo repositories.IFormRepo, formVersionRepo repositories.IFormVersionRepo) *SubmissionService {
	return &SubmissionService{
		formRepo:        formRepo,
		formVersionRepo: formVersionRepo,
	}
}

func (s *SubmissionService) GetForm(ctx context.Context, domain string) (*types.GetSubmissionFormResponse, error) {

	form, err := s.formRepo.GetByDomain(ctx, domain)
	if err != nil {
		return nil, err
	}

	if form == nil {
		return nil, errors.NotFound("Form not found")
	}

	if form.UniqueSubdomain == nil && form.CustomDomain == nil {
		return nil, errors.BadRequest("Form does not have a publish URL")
	}

	version, err := s.formVersionRepo.GetLatestVersion(ctx, form.ID)
	if err != nil {
		return nil, err
	}
	if version == nil {
		return nil, errors.NotFound("Published version not found")
	}

	var snapshot types.PublishVersionSnapshot
	err = json.Unmarshal([]byte(version.Schema), &snapshot)
	if err != nil {
		return nil, errors.Internal(err)
	}

	resp := &types.GetSubmissionFormResponse{
		FormID:      form.ID,
		PublishedAt: utils.GetIsoDateTime(version.PublishedAt),
		FormVersion: version.VersionNumber,
		FormMetadata: &types.SubmissionFormMetadata{
			Name:                form.Name,
			Description:         form.Description,
			Theme:               form.Theme,
			LogoURL:             form.LogoURL,
			RequireConsent:      form.RequireConsent,
			AllowAnonymous:      form.AllowAnonymous,
			MultipleSubmissions: form.MultipleSubmissions,
			PasswordProtected:   form.PasswordProtected,
			FormPageType:        string(form.FormPageType),
			Metadata:            form.Metadata,
		},
		Questions: snapshot.Questions,
		Edges:     snapshot.Edges,
	}

	return resp, nil
}
