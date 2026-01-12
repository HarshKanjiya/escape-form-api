package models

import "time"

type Feature struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	PlanID      string    `gorm:"type:uuid;column:plan_id" json:"planId"`
	Key         string    `gorm:"column:key" json:"key"`
	Name        *string   `gorm:"column:name" json:"name"`
	Description *string   `gorm:"column:description" json:"description"`
	Valid       bool      `gorm:"default:true;column:valid" json:"valid"`
	CreatedAt   time.Time `gorm:"type:timestamptz(6);default:now();column:created_at" json:"createdAt"`
	UpdatedAt   *time.Time `gorm:"type:timestamptz(6);column:updated_at" json:"updatedAt"`
	Plan        Plan      `gorm:"references:ID" json:"plan"`
}

func (Feature) TableName() string {
	return "features"
}