package types

import "github.com/HarshKanjiya/escape-form-api/internal/models"

// Response structs
type FormResponse struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	TeamID              string `json:"teamId"`
	ProjectID           string `json:"projectId"`
	Theme               string `json:"theme"`
	LogoURL             string `json:"logoUrl"`
	MaxResponses        *int   `json:"maxResponses"`
	OpenAt              string `json:"openAt"`
	CloseAt             string `json:"closeAt"`
	Status              string `json:"status"`
	UniqueSubdomain     string `json:"uniqueSubdomain"`
	CustomDomain        string `json:"customDomain"`
	RequireConsent      *bool  `json:"requireConsent"`
	AllowAnonymous      *bool  `json:"allowAnonymous"`
	MultipleSubmissions *bool  `json:"multipleSubmissions"`
	PasswordProtected   *bool  `json:"passwordProtected"`
	AnalyticsEnabled    *bool  `json:"analyticsEnabled"`
	Valid               bool   `json:"valid"`
	Metadata            string `json:"metadata"`
	CreatedBy           string `json:"createdBy"`
	CreatedAt           string `json:"createdAt"`
	UpdatedAt           string `json:"updatedAt"`
	FormPageType        string `json:"formPageType"`
	ResponseCount       int    `json:"responseCount,omitempty"`
	Questions           []any  `json:"questions"`
	Edges               []any  `json:"edges"`
}

type CreateFormRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	ProjectID   string  `json:"projectId"`
}

type UpdateFormStatusRequest struct {
	Status models.FormStatus `json:"status" validate:"required,oneof=DRAFT OPEN CLOSED"`
}

type UpdateSequenceRequest struct {
	Sequence []*SequenceItem `json:"sequence"`
}

type SequenceItem struct {
	ID       string `json:"id"`
	NewOrder int    `json:"newOrder"`
}
