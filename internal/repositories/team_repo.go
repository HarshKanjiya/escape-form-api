package repositories

import (
	"log"

	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
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

func (r *TeamRepo) Get(ctx *fiber.Ctx, pagination *types.PaginationQuery, valid bool) ([]*types.TeamResponse, error) {

	t := r.q.Team
	userId := ctx.Locals("user_id").(string)
	query := r.q.
		WithContext(ctx.Context()).
		Team.Where(t.OwnerID.Eq(userId), t.Valid.Is(valid))

	if pagination.Search != "" {
		query = query.Where(t.Name.Like("%" + pagination.Search + "%"))
	}

	query.Limit(pagination.Limit)
	query.Offset((pagination.Page - 1) * pagination.Limit)

	teams, err := query.Find()
	if err != nil {
		log.Printf("Error fetching teams: %v", err)
		return nil, err
	}

	if len(teams) == 0 {
		log.Printf("No teams found for user %s", userId)
		return []*types.TeamResponse{}, nil
	}

	teamIDs := make([]string, len(teams))
	for i, team := range teams {
		teamIDs[i] = team.ID
	}

	projectCounts := make(map[string]int)
	var results []struct {
		TeamID string
		Count  int
	}
	err = r.q.WithContext(ctx.Context()).
		Project.Select(r.q.Project.TeamID, r.q.Project.ID.Count().As("count")).
		Where(r.q.Project.TeamID.In(teamIDs...)).
		Group(r.q.Project.TeamID).
		Scan(&results)
	if err != nil {
		log.Printf("Error fetching project counts: %v", err)
		// Continue without counts (set to 0)
	} else {
		for _, res := range results {
			projectCounts[res.TeamID] = res.Count
		}
	}

	var teamResponses []*types.TeamResponse
	for _, team := range teams {
		name := ""
		if team.Name != nil {
			name = *team.Name
		}
		ownerId := ""
		if team.OwnerID != nil {
			ownerId = *team.OwnerID
		}
		planId := ""
		if team.PlanID != nil {
			planId = *team.PlanID
		}
		teamResponses = append(teamResponses, &types.TeamResponse{
			ID:           team.ID,
			Name:         name,
			OwnerId:      ownerId,
			PlanId:       planId,
			Valid:        team.Valid,
			ProjectCount: projectCounts[team.ID], // Use actual count
			CreatedAt:    utils.GetIsoDateTime(&team.CreatedAt),
			UpdatedAt:    utils.GetIsoDateTime(team.UpdatedAt),
		})
	}

	return teamResponses, nil
}
