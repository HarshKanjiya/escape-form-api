package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
)

type DashService struct {
	dashRepo *repositories.DashRepo
}

func NewDashService(dashRepo *repositories.DashRepo) *DashService {
	return &DashService{
		dashRepo: dashRepo,
	}
}
