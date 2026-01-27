package types

// Request structs
type CreateTeamRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}

type GetTeamsRequest struct {
	PaginationQuery
}

type UpdateTeamRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}

// Response structs
type TeamResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	OwnerId      string `json:"ownerId"`
	PlanId       string `json:"planId"`
	Valid        bool   `json:"valid"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
	ProjectCount int    `json:"projectCount,omitempty"`
}

type TeamDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
