package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/gofiber/fiber/v2"
)

type IEdgeService interface {
}

type EdgeService struct {
	edgeRepo repositories.IEdgeRepo
}

func NewEdgeService(edgeRepo repositories.IEdgeRepo) *EdgeService {
	return &EdgeService{
		edgeRepo: edgeRepo,
	}
}

func (s *EdgeService) Get(ctx *fiber.Ctx, formId string) ([]*models.Edge, error) {
	return s.edgeRepo.Get(ctx, formId)
}

func (s *EdgeService) Create(ctx *fiber.Ctx, edge *types.EdgeDto) (*models.Edge, error) {
	return s.edgeRepo.Create(ctx, edge)
}

func (s *EdgeService) Update(ctx *fiber.Ctx, edge *types.EdgeDto) (*models.Edge, error) {
	return s.edgeRepo.Update(ctx, edge)
}

func (s *EdgeService) Delete(ctx *fiber.Ctx, edgeId string) error {
	return s.edgeRepo.Delete(ctx, edgeId)
}
