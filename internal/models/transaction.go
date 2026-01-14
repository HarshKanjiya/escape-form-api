package models

import "time"

type Transaction struct {
	ID          string          `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	Type        TransactionType `gorm:"column:type" json:"type"`
	Amount      float64         `gorm:"column:amount" json:"amount"`
	Description *string         `gorm:"column:description" json:"description"`
	CreatedAt   time.Time       `gorm:"type:timestamptz(6);default:now();column:createdAt" json:"createdAt"`
	CreatedBy   *string         `gorm:"type:varchar;column:createdBy" json:"createdBy"`
	TeamID      string          `gorm:"type:uuid;index;column:teamId" json:"teamId"`
	Team        Team            `gorm:"references:ID" json:"team"`
}

func (Transaction) TableName() string {
	return "transactions"
}