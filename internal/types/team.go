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
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
