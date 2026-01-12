package controllers

import (
	"github.com/HarshKanjiya/escape-form-api/internal/config"
	"github.com/HarshKanjiya/escape-form-api/internal/services"
)

type TeamController struct {
}

func NewTeamController(*services.TeamService, *config.Config) *TeamController {
	return &TeamController{}
}
