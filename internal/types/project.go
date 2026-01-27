package types

// Request structs
type CreateProjectRequest struct {
	Name        string  `json:"name" validate:"required,min=3,max=100"`
	Description *string `json:"description"`
	TeamID      string  `json:"teamId" validate:"required"`
}

type GetProjectsRequest struct {
	PaginationQuery
}

type UpdateProjectRequest struct {
	Name        string  `json:"name" validate:"required,min=3,max=100"`
	Description *string `json:"description"`
}

// Response structs
type ProjectResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TeamID      string `json:"teamId"`
	Valid       bool   `json:"valid"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	FormCount   int    `json:"formCount,omitempty"`
}

type ProjectDto struct {
	ID          string  `json:"id"`
	Name        string  `json:"name" validate:"required,min=3,max=100"`
	Description *string `json:"description"`
	TeamID      string  `json:"teamId" validate:"required"`
}
