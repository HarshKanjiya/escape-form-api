package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"gorm.io/gorm"
)

type DashService struct {
	q *query.Query
}

func NewDashService(db *gorm.DB) *DashService {
	return &DashService{
	q: query.Use(db),
	}
}
