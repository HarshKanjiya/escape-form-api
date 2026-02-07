package models

import "time"

type UsageEvents struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	TeamID    string    `gorm:"type:uuid;index;column:teamId" json:"teamId"`
	Type      string    `gorm:"column:type" json:"type"`
	Quantity  int       `gorm:"column:quantity" json:"quantity"`
	RefID     *string   `gorm:"type:uuid;column:refId" json:"refId"`
	CreatedAt time.Time `gorm:"type:timestamptz(6);default:now();column:createdAt" json:"createdAt"`
}

func (UsageEvents) TableName() string {
	return "usage_events"
}
