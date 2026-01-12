package types

type CreateTeamRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}

type GetTeamsRequest struct {
	Pagination
}

type UpdateTeamRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}
