package models

import "time"

type Project struct {
	ID          string     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	Name        string     `gorm:"column:name" json:"name"`
	Description *string    `gorm:"column:description" json:"description"`
	TeamID      string     `gorm:"type:uuid;index;column:teamId" json:"teamId"`
	Valid       bool       `gorm:"default:true;column:valid" json:"valid"`
	CreatedAt   *time.Time `gorm:"type:timestamptz(6);column:createdAt" json:"createdAt"`
	UpdatedAt   *time.Time `gorm:"type:timestamp(6);column:updatedAt" json:"updatedAt"`
	Forms       []Form     `gorm:"foreignKey:ProjectID" json:"forms"`
	Team        Team       `gorm:"references:ID" json:"team"`
}

func (Project) TableName() string {
	return "projects"
}
