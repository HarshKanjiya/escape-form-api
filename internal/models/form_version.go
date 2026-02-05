package models

import (
	"time"

	"gorm.io/datatypes"
)

type FormVersion struct {
	ID            string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	FormID        string         `gorm:"type:uuid;index;column:formId" json:"formId"`
	VersionNumber int            `gorm:"default:1;column:versionNumber" json:"versionNumber"`
	Schema        datatypes.JSON `gorm:"type:jsonb;column:schema" json:"schema"`
	CreatedAt     time.Time      `gorm:"type:timestamptz(6);default:CURRENT_TIMESTAMP;column:createdAt" json:"createdAt"`
	PublishedAt   *time.Time     `gorm:"type:timestamptz(6);column:publishedAt" json:"publishedAt"`
	Form          Form           `gorm:"foreignKey:FormID;references:ID;onDelete:CASCADE" json:"form"`
}

func (FormVersion) TableName() string {
	return "form_versions"
}
