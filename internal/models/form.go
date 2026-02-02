package models

import (
	"time"

	"gorm.io/datatypes"
)

type Form struct {
	ID                  string           `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	Name                string           `gorm:"column:name" json:"name"`
	Description         *string          `gorm:"column:description" json:"description"`
	TeamID              string           `gorm:"type:uuid;index;column:teamId" json:"teamId"`
	ProjectID           string           `gorm:"type:uuid;index;column:projectId" json:"projectId"`
	Theme               *string          `gorm:"column:theme" json:"theme"`
	LogoURL             *string          `gorm:"column:logoUrl" json:"logoUrl"`
	MaxResponses        *int             `gorm:"column:maxResponses" json:"maxResponses"`
	OpenAt              *time.Time       `gorm:"type:timestamptz(6);column:openAt" json:"openAt"`
	CloseAt             *time.Time       `gorm:"type:timestamptz(6);column:closeAt" json:"closeAt"`
	Status              *FormStatus      `gorm:"column:status" json:"status"`
	UniqueSubdomain     *string          `gorm:"column:uniqueSubdomain" json:"uniqueSubdomain"`
	CustomDomain        *string          `gorm:"column:customDomain" json:"customDomain"`
	RequireConsent      *bool            `gorm:"column:requireConsent" json:"requireConsent"`
	AllowAnonymous      *bool            `gorm:"column:allowAnonymous" json:"allowAnonymous"`
	MultipleSubmissions *bool            `gorm:"default:false;column:multipleSubmissions" json:"multipleSubmissions"`
	PasswordProtected   *bool            `gorm:"default:false;column:passwordProtected" json:"passwordProtected"`
	AnalyticsEnabled    *bool            `gorm:"default:true;column:analyticsEnabled" json:"analyticsEnabled"`
	Valid               bool             `gorm:"default:true;column:valid" json:"valid"`
	Metadata            datatypes.JSON   `gorm:"type:jsonb;default:'{}';column:metadata" json:"metadata"`
	CreatedBy           string           `gorm:"column:createdBy" json:"createdBy"`
	CreatedAt           *time.Time       `gorm:"type:timestamptz(6);column:createdAt" json:"createdAt"`
	UpdatedAt           *time.Time       `gorm:"type:timestamp(6);autoUpdateTime;column:updatedAt" json:"updatedAt"`
	FormPageType        FormPageType     `gorm:"default:'STEPPER';column:formPageType" json:"formPageType"`
	ActivePasswords     []ActivePassword `gorm:"foreignKey:FormID" json:"activePasswords"`
	Edges               []Edge           `gorm:"foreignKey:FormID" json:"edges"`
	Project             Project          `gorm:"foreignKey:ProjectID;references:ID;onDelete:CASCADE" json:"project"`
	Team                Team             `gorm:"foreignKey:TeamID;references:ID;onDelete:CASCADE" json:"team"`
	Questions           []Question       `gorm:"foreignKey:FormID" json:"questions"`
	Responses           []Response       `gorm:"foreignKey:FormID" json:"responses"`
	ResponseCount       int              `gorm:"-" json:"responseCount,omitempty"`
}

func (Form) TableName() string {
	return "forms"
}
