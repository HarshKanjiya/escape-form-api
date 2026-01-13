package repositories

import (
	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TeamRepo struct {
	q *query.Query
}

func NewTeamRepo(db *gorm.DB) *TeamRepo {
	return &TeamRepo{
		q: query.Use(db),
	}
}

func (r *TeamRepo) Get(ctx *fiber.Ctx, pagination *types.PaginationQuery, valid bool) ([]*models.Team, error) {

	t := r.q.Team
	query := r.q.WithContext(ctx.Context()).Team.Where(t.OwnerID.Eq("123"), t.Valid.Is(valid))

	if pagination.Search != "" {
		query = query.Where(t.Name.Like("%" + pagination.Search + "%"))
	}

	query.Limit(pagination.Limit)
	query.Offset((pagination.Page - 1) * pagination.Limit)

	teams, err := query.Find()
	if err != nil {
		return nil, err
	}

	return teams, nil
}
