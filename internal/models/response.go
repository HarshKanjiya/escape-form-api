package models

import "time"

type Response struct {
	ID          string          `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	FormID      string          `gorm:"type:uuid;index;column:formId" json:"formId"`
	UserID      *string         `gorm:"type:uuid;column:userId" json:"userId"`
	Data        interface{}     `gorm:"type:jsonb;default:'{}';column:data" json:"data"`
	MetaData    *interface{}    `gorm:"type:jsonb;default:'{}';column:metaData" json:"metaData"`
	Tags        []string        `gorm:"type:text[];column:tags" json:"tags"`
	Status      *ResponseStatus `gorm:"column:status" json:"status"`
	PartialSave *bool           `gorm:"column:partialSave" json:"partialSave"`
	Notified    *bool           `gorm:"column:notified" json:"notified"`
	Valid       bool            `gorm:"default:true;column:valid" json:"valid"`
	StartedAt   *time.Time      `gorm:"type:timestamptz(6);column:startedAt" json:"startedAt"`
	SubmittedAt *time.Time      `gorm:"type:timestamptz(6);column:submittedAt" json:"submittedAt"`
	UpdatedAt   *time.Time      `gorm:"type:timestamp(6);autoUpdateTime;column:updatedAt" json:"updatedAt"`
	Form        Form            `gorm:"foreignKey:FormID;references:ID;onDelete:CASCADE" json:"form"`
}

func (Response) TableName() string {
	return "responses"
}
