package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
)

type EdgeService struct {
	edgeRepo *repositories.EdgeRepo
}

func NewEdgeService(edgeRepo *repositories.EdgeRepo) *EdgeService {
	return &EdgeService{
		edgeRepo: edgeRepo,
	}
}
