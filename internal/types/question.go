package types

type QuestionOptionDto struct {
	ID         string `json:"id"`
	QuestionID string `json:"questionId" validate:"required"`
	Label      string `json:"label" validate:"required"`
	Value      string `json:"value" validate:"required"`
	SortOrder  int    `json:"sortOrder"`
}