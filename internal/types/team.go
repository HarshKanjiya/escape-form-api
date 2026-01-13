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
	ID        string `json:"id"`
	Name      string `json:"name"`
	OwnerId   string `json:"owner_id"`
	PlanId    string `json:"plan_id"`
	Valid     bool   `json:"valid"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
