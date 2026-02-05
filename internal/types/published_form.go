package types

import "gorm.io/datatypes"

type PublishVersionSnapshot struct {
	Name                string                       `json:"name"`
	Description         *string                      `json:"description"`
	Theme               datatypes.JSON               `json:"theme"`
	LogoURL             *string                      `json:"logoUrl"`
	RequireConsent      *bool                        `json:"requireConsent"`
	AllowAnonymous      *bool                        `json:"allowAnonymous"`
	MultipleSubmissions *bool                        `json:"multipleSubmissions"`
	PasswordProtected   *bool                        `json:"passwordProtected"`
	FormPageType        string                       `json:"formPageType"`
	Metadata            datatypes.JSON               `json:"metadata"`
	Questions           []PublishedQuestion          `json:"questions"`
	Edges               []PublishedEdge              `json:"edges"`
}

type PublishedQuestion struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	Placeholder string                 `json:"placeholder"`
	Description string                 `json:"description"`
	Required    bool                   `json:"required"`
	Type        string                 `json:"type"`
	Metadata    datatypes.JSON         `json:"metadata"`
	PosX        int                    `json:"posX"`
	PosY        int                    `json:"posY"`
	SortOrder   *int                   `json:"sortOrder"`
	Options     []PublishedQuestionOption `json:"options"`
}

type PublishedQuestionOption struct {
	ID        string `json:"id"`
	Label     string `json:"label"`
	Value     string `json:"value"`
	SortOrder int    `json:"sortOrder"`
}

type PublishedEdge struct {
	ID           string      `json:"id"`
	SourceNodeID string      `json:"sourceNodeId"`
	TargetNodeID string      `json:"targetNodeId"`
	Condition    interface{} `json:"condition"`
}
