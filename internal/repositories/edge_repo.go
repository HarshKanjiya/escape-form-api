package repositories

import (
	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EdgeRepo struct {
	q *query.Query
}

func NewEdgeRepo(db *gorm.DB) *EdgeRepo {
	return &EdgeRepo{
		q: query.Use(db),
	}
}

func (r *EdgeRepo) Get(ctx *fiber.Ctx, formId string) ([]*models.Edge, error) {
	edges, err := r.q.WithContext(ctx.Context()).
		Edge.Where(r.q.Edge.FormID.Eq(formId)).Find()

	if err != nil {
		return nil, err
	}
	return edges, nil
}

func (r *EdgeRepo) Create(ctx *fiber.Ctx, edge *types.EdgeDto) (*models.Edge, error) {

	edgeModel := &models.Edge{
		ID:           uuid.New().String(),
		FormID:       edge.FormID,
		SourceNodeID: edge.SourceNodeID,
		TargetNodeID: edge.TargetNodeID,
		Condition:    edge.Condition,
	}

	if err := r.q.Edge.Create(edgeModel); err != nil {
		return nil, err
	}
	return edgeModel, nil
}

func (r *EdgeRepo) Update(ctx *fiber.Ctx, edge *types.EdgeDto) (*models.Edge, error) {
	edgeModel := &models.Edge{
		ID:           edge.ID,
		FormID:       edge.FormID,
		SourceNodeID: edge.SourceNodeID,
		TargetNodeID: edge.TargetNodeID,
		Condition:    edge.Condition,
	}
	_, err := r.q.WithContext(ctx.Context()).
		Edge.Where(r.q.Edge.ID.Eq(edge.ID)).
		Updates(edgeModel)

	if err != nil {
		return nil, err
	}

	return edgeModel, nil
}

func (r *EdgeRepo) Delete(ctx *fiber.Ctx, edgeId string) error {

	_, err := r.q.WithContext(ctx.Context()).
		Edge.Where(r.q.Edge.ID.Eq(edgeId)).
		Delete()
	if err != nil {
		return err
	}
	return nil
}
