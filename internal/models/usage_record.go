package models

import "time"

type UsageRecord struct {
	ID            string          `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	TeamID        string          `gorm:"type:uuid;index;column:teamId" json:"teamId"`
	CycleStart    time.Time       `gorm:"type:timestamptz(6);column:cycleStart" json:"cycleStart"`
	FormsUsed     int             `gorm:"default:0;column:formsUsed" json:"formsUsed"`
	ProjectsUsed  int             `gorm:"default:0;column:projectsUsed" json:"projectsUsed"`
	ResponsesUsed int             `gorm:"default:0;column:responsesUsed" json:"responsesUsed"`
	StorageUsed   int             `gorm:"default:0;column:storageUsed" json:"storageUsed"`
	TokensUsed    int             `gorm:"default:0;column:tokensUsed" json:"tokensUsed"`
	Type          TransactionType `gorm:"column:type" json:"type"`
	Amount        float64         `gorm:"column:amount" json:"amount"`
	RecordedAt    time.Time       `gorm:"type:timestamptz(6);default:now();column:recordedAt" json:"recordedAt"`
	Team          Team            `gorm:"foreignKey:TeamID;references:ID;onDelete:CASCADE" json:"team"`
}

func (UsageRecord) TableName() string {
	return "usage_records"
}
