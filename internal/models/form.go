package models

import "time"

type Form struct {
	ID                  string           `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	Name                string           `gorm:"column:name" json:"name"`
	Description         *string          `gorm:"column:description" json:"description"`
	TeamID              string           `gorm:"type:uuid;index;column:team_id" json:"teamId"`
	ProjectID           string           `gorm:"type:uuid;index;column:project_id" json:"projectId"`
	Theme               *string          `gorm:"column:theme" json:"theme"`
	LogoURL             *string          `gorm:"column:logo_url" json:"logoUrl"`
	MaxResponses        *int             `gorm:"column:max_responses" json:"maxResponses"`
	OpenAt              *time.Time       `gorm:"type:timestamptz(6);column:open_at" json:"openAt"`
	CloseAt             *time.Time       `gorm:"type:timestamptz(6);column:close_at" json:"closeAt"`
	Status              *FormStatus      `gorm:"column:status" json:"status"`
	UniqueSubdomain     *string          `gorm:"column:unique_subdomain" json:"uniqueSubdomain"`
	CustomDomain        *string          `gorm:"column:custom_domain" json:"customDomain"`
	RequireConsent      *bool            `gorm:"column:require_consent" json:"requireConsent"`
	AllowAnonymous      *bool            `gorm:"column:allow_anonymous" json:"allowAnonymous"`
	MultipleSubmissions *bool            `gorm:"default:false;column:multiple_submissions" json:"multipleSubmissions"`
	PasswordProtected   *bool            `gorm:"default:false;column:password_protected" json:"passwordProtected"`
	AnalyticsEnabled    *bool            `gorm:"default:true;column:analytics_enabled" json:"analyticsEnabled"`
	Valid               bool             `gorm:"default:true;column:valid" json:"valid"`
	Metadata            *string          `gorm:"type:jsonb;default:'{}';column:metadata" json:"metadata"`
	CreatedBy           string           `gorm:"column:created_by" json:"createdBy"`
	CreatedAt           *time.Time       `gorm:"type:timestamptz(6);column:created_at" json:"createdAt"`
	UpdatedAt           *time.Time       `gorm:"type:timestamp(6);column:updated_at" json:"updatedAt"`
	FormPageType        FormPageType     `gorm:"default:'STEPPER';column:form_page_type" json:"formPageType"`
	ActivePasswords     []ActivePassword `gorm:"foreignKey:FormID" json:"activePasswords"`
	Edges               []Edge           `gorm:"foreignKey:FormID" json:"edges"`
	Project             Project          `gorm:"references:ID" json:"project"`
	Team                Team             `gorm:"references:ID" json:"team"`
	Questions           []Question       `gorm:"foreignKey:FormID" json:"questions"`
	Responses           []Response       `gorm:"foreignKey:FormID" json:"responses"`
}

func (Form) TableName() string {
	return "forms"
}
