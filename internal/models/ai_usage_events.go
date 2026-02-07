package models

import "time"

type AiUsageEvents struct {
	ID            string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	TeamID        string    `gorm:"type:uuid;index;column:teamId" json:"teamId"`
	Model         string    `gorm:"column:model" json:"model"`
	Feature       *string   `gorm:"column:feature" json:"feature"`
	InputTokens   int       `gorm:"column:inputTokens" json:"inputTokens"`
	OutputTokens  int       `gorm:"column:outputTokens" json:"outputTokens"`
	EstimatedCost float64   `gorm:"column:estimatedCost" json:"estimatedCost"`
	RefID         *string   `gorm:"type:uuid;column:refId" json:"refId"`
	CreatedAt     time.Time `gorm:"type:timestamptz(6);default:now();column:createdAt" json:"createdAt"`
}

func (AiUsageEvents) TableName() string {
	return "ai_usage_events"
}
