package types

import "gorm.io/datatypes"

type GetSubmissionFormResponse struct {
	FormID      string `json:"formId"`
	FormVersion int    `json:"formVersion"`
	PublishedAt string `json:"publishedAt"`

	FormMetadata *SubmissionFormMetadata `json:"formMetadata"`
	Questions    []PublishedQuestion     `json:"questions"`
	Edges        []PublishedEdge         `json:"edges"`
}

type SubmissionFormMetadata struct {
	Name                string         `json:"name"`
	Description         *string        `json:"description"`
	Theme               datatypes.JSON `json:"theme"`
	LogoURL             *string        `json:"logoUrl"`
	RequireConsent      *bool          `json:"requireConsent"`
	AllowAnonymous      *bool          `json:"allowAnonymous"`
	MultipleSubmissions *bool          `json:"multipleSubmissions"`
	PasswordProtected   *bool          `json:"passwordProtected"`
	FormPageType        string         `json:"formPageType"`
	Metadata            datatypes.JSON `json:"metadata"`
}
