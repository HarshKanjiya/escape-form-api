package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
)

type FormService struct {
	formRepo *repositories.FormRepo
}

func NewFormService(formRepo *repositories.FormRepo) *FormService {
	return &FormService{
		formRepo: formRepo,
	}
}
