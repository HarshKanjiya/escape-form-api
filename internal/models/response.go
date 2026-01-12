package models

import "time"

type Response struct {
	ID          string          `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	FormID      string          `gorm:"type:uuid;index;column:form_id" json:"formId"`
	UserID      *string         `gorm:"type:uuid;column:user_id" json:"userId"`
	Data        string          `gorm:"type:jsonb;default:'{}';column:data" json:"data"`
	MetaData    *string         `gorm:"type:jsonb;default:'{}';column:meta_data" json:"metaData"`
	Tags        []string        `gorm:"type:text[];column:tags" json:"tags"`
	Status      *ResponseStatus `gorm:"column:status" json:"status"`
	PartialSave *bool           `gorm:"column:partial_save" json:"partialSave"`
	Notified    *bool           `gorm:"column:notified" json:"notified"`
	Valid       bool            `gorm:"default:true;column:valid" json:"valid"`
	StartedAt   *time.Time      `gorm:"type:timestamptz(6);column:started_at" json:"startedAt"`
	SubmittedAt *time.Time      `gorm:"type:timestamptz(6);column:submitted_at" json:"submittedAt"`
	UpdatedAt   *time.Time      `gorm:"type:timestamp(6);column:updated_at" json:"updatedAt"`
	Form        Form            `gorm:"references:ID" json:"form"`
}

func (Response) TableName() string {
	return "responses"
}
