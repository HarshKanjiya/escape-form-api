package types

import (
	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"gorm.io/datatypes"
)

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
	Metadata    datatypes.JSON      `json:"metadata"`
	PosX        int                 `json:"posX"`
	PosY        int                 `json:"posY"`
	SortOrder   *int                `json:"sortOrder"`
}

type QuestionResponse struct {
	ID            string               `json:"id"`
	FormID        string               `json:"formId"`
	Title         string               `json:"title"`
	Placeholder   string               `json:"placeholder"`
	Description   string               `json:"description"`
	Required      bool                 `json:"required"`
	Type          models.QuestionType  `json:"type"`
	Metadata      datatypes.JSON       `json:"metadata"`
	PosX          int                  `json:"posX"`
	PosY          int                  `json:"posY"`
	SortOrder     *int                 `json:"sortOrder"`
	OutgoingEdges []*EdgeResponse      `json:"outgoingEdges"`
	IncomingEdges []*EdgeResponse      `json:"incomingEdges"`
	Options       []*QueOptionResponse `json:"options"`
}

type QueOptionResponse struct {
	ID         string `json:"id"`
	QuestionID string `json:"questionId"`
	Label      string `json:"label"`
	Value      string `json:"value"`
	SortOrder  int    `json:"sortOrder"`
}
