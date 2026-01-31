package types

import "github.com/HarshKanjiya/escape-form-api/internal/models"

type QuestionOptionRequest struct {
	ID         string `json:"id"`
	QuestionID string `json:"questionId" validate:"required"`
	Label      string `json:"label" validate:"required"`
	Value      string `json:"value" validate:"required"`
	SortOrder  int    `json:"sortOrder"`
}

type QuestionRequest struct {
	ID          string              `json:"id"`
	FormID      string              `json:"formId"`
	Title       string              `json:"title"`
	Placeholder string              `json:"placeholder"`
	Description string              `json:"description"`
	Required    bool                `json:"required"`
	Type        models.QuestionType `json:"type"`
	Metadata    interface{}         `json:"metadata"`
	PosX        int                 `json:"posX"`
	PosY        int                 `json:"posY"`
	SortOrder   *int                `json:"sortOrder"`
}
