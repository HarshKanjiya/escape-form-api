package models

import "time"

type Team struct {
	ID           string        `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	Name         *string       `gorm:"type:varchar;column:name" json:"name"`
	OwnerID      *string       `gorm:"type:varchar;index;column:owner_id" json:"ownerId"`
	PlanID       *string       `gorm:"type:uuid;column:plan_id" json:"planId"`
	Valid        bool          `gorm:"default:true;column:valid" json:"valid"`
	CreatedAt    time.Time     `gorm:"type:timestamptz(6);default:now();column:created_at" json:"createdAt"`
	UpdatedAt    *time.Time    `gorm:"type:timestamp(6);column:updated_at" json:"updatedAt"`
	Forms        []Form        `gorm:"foreignKey:TeamID" json:"forms"`
	Projects     []Project     `gorm:"foreignKey:TeamID" json:"projects"`
	Plan         *Plan         `gorm:"foreignKey:PlanID" json:"plan"`
	Transactions []Transaction `gorm:"foreignKey:TeamID" json:"transactions"`
}

func (Team) TableName() string {
	return "teams"
}
