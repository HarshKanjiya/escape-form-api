package types

// Request structs
type CreateFormRequest struct {
	Name                string  `json:"name" validate:"required,min=3,max=100"`
	Description         *string `json:"description"`
	ProjectID           string  `json:"projectId" validate:"required"`
	Theme               *string `json:"theme"`
	LogoURL             *string `json:"logoUrl"`
	MaxResponses        *int    `json:"maxResponses"`
	OpenAt              *string `json:"openAt"`
	CloseAt             *string `json:"closeAt"`
	Status              *string `json:"status"`
	UniqueSubdomain     *string `json:"uniqueSubdomain"`
	CustomDomain        *string `json:"customDomain"`
	RequireConsent      *bool   `json:"requireConsent"`
	AllowAnonymous      *bool   `json:"allowAnonymous"`
	MultipleSubmissions *bool   `json:"multipleSubmissions"`
	PasswordProtected   *bool   `json:"passwordProtected"`
	AnalyticsEnabled    *bool   `json:"analyticsEnabled"`
	Metadata            *string `json:"metadata"`
	FormPageType        *string `json:"formPageType"`
}

type GetFormsRequest struct {
	PaginationQuery
}

type UpdateFormRequest struct {
	Name                string  `json:"name" validate:"required,min=3,max=100"`
	Description         *string `json:"description"`
	Theme               *string `json:"theme"`
	LogoURL             *string `json:"logoUrl"`
	MaxResponses        *int    `json:"maxResponses"`
	OpenAt              *string `json:"openAt"`
	CloseAt             *string `json:"closeAt"`
	Status              *string `json:"status"`
	UniqueSubdomain     *string `json:"uniqueSubdomain"`
	CustomDomain        *string `json:"customDomain"`
	RequireConsent      *bool   `json:"requireConsent"`
	AllowAnonymous      *bool   `json:"allowAnonymous"`
	MultipleSubmissions *bool   `json:"multipleSubmissions"`
	PasswordProtected   *bool   `json:"passwordProtected"`
	AnalyticsEnabled    *bool   `json:"analyticsEnabled"`
	Metadata            *string `json:"metadata"`
	FormPageType        *string `json:"formPageType"`
}

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
	Questions           []any  `json:"questions,omitempty"`
	Edges               []any  `json:"edges,omitempty"`
}
