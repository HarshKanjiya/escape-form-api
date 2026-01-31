package services

import (
	"context"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
)

type IEdgeService interface {
	Get(ctx context.Context, userId string, formId string) ([]*models.Edge, error)
	Create(ctx context.Context, userId string, formId string, edge *types.CreateEdgeRequest) (*models.Edge, error)
	Update(ctx context.Context, userId string, formId string, edgeId string, edge *types.UpdateEdgeRequest) error
	Delete(ctx context.Context, userId string, formId string, edgeId string) error
}

type EdgeService struct {
	edgeRepo repositories.IEdgeRepo
	formRepo repositories.IFormRepo
}

func NewEdgeService(edgeRepo repositories.IEdgeRepo, formRepo repositories.IFormRepo) *EdgeService {
	return &EdgeService{
		edgeRepo: edgeRepo,
		formRepo: formRepo,
	}
}

func (s *EdgeService) Get(ctx context.Context, userId string, formId string) ([]*models.Edge, error) {

	form, err := s.formRepo.GetByIdWithTeam(ctx, formId)
	if err != nil {
		return nil, err
	}

	if form == nil {
		return nil, errors.NotFound("Form")
	}

	if form.Team.OwnerID == nil || *form.Team.OwnerID != userId {
		return nil, errors.Unauthorized("")
	}

	edges, err := s.edgeRepo.Get(ctx, formId)
	if err != nil {
		return nil, err
	}
	return edges, nil
}

func (s *EdgeService) Create(ctx context.Context, userId string, formId string, edge *types.CreateEdgeRequest) (*models.Edge, error) {

	form, err := s.formRepo.GetByIdWithTeam(ctx, formId)
	if err != nil {
		return nil, err
	}
	if form == nil {
		return nil, errors.NotFound("Form")
	}
	if form.Team.OwnerID == nil || *form.Team.OwnerID != userId {
		return nil, errors.Unauthorized("")
	}
	newEdge := &models.Edge{
		ID:           utils.GenerateUUID(),
		FormID:       formId,
		SourceNodeID: edge.SourceNodeID,
		TargetNodeID: edge.TargetNodeID,
		Condition:    nil,
	}
	createdEdge, err := s.edgeRepo.Create(ctx, newEdge)
	if err != nil {
		return nil, err
	}
	return createdEdge, nil
}

func (s *EdgeService) Update(ctx context.Context, userId string, formId string, edgeId string, edge *types.UpdateEdgeRequest) error {

	form, err := s.formRepo.GetByIdWithTeam(ctx, formId)
	if err != nil {
		return err
	}
	if form == nil {
		return errors.NotFound("Form")
	}
	if form.Team.OwnerID == nil || *form.Team.OwnerID != userId {
		return errors.Unauthorized("")
	}

	edgeModel := &models.Edge{
		ID:        edgeId,
		Condition: edge.Condition,
	}

	err = s.edgeRepo.Update(ctx, edgeId, edgeModel)
	if err != nil {
		return err
	}
	return nil
}

func (s *EdgeService) Delete(ctx context.Context, userId string, formId string, edgeId string) error {

	form, err := s.formRepo.GetByIdWithTeam(ctx, formId)
	if err != nil {
		return err
	}
	if form == nil {
		return errors.NotFound("Form")
	}
	if form.Team.OwnerID == nil || *form.Team.OwnerID != userId {
		return errors.Unauthorized("")
	}
	err = s.edgeRepo.Delete(ctx, edgeId)
	if err != nil {
		return err
	}
	return nil
}
