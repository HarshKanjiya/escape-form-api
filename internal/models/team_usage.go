package models

import "time"

type TeamUsage struct {
	ID              string          `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	TeamID          string          `gorm:"type:uuid;column:teamId" json:"teamId"`
	FormsAllowed    int             `gorm:"default:1;column:formsAllowed" json:"formsAllowed"`
	ProjectsAllowed int             `gorm:"default:1;column:projectsAllowed" json:"projectsAllowed"`
	Responses       int             `gorm:"default:0;column:responses" json:"responses"`
	OverUsage       int             `gorm:"default:0;column:overUsage" json:"overUsage"`
	CycleStart      time.Time       `gorm:"type:timestamptz(6);column:cycleStart" json:"cycleStart"`
	CycleEnd        time.Time       `gorm:"type:timestamptz(6);column:cycleEnd" json:"cycleEnd"`
	Status          TeamUsageStatus `gorm:"default:ACTIVE;column:status" json:"status"`
	CreatedAt       time.Time       `gorm:"type:timestamptz(6);default:now();column:createdAt" json:"createdAt"`
	UpdatedAt       *time.Time      `gorm:"type:timestamp(6);autoUpdateTime;column:updatedAt" json:"updatedAt"`
	Team            *Team           `gorm:"foreignKey:TeamID" json:"team"`
}

func (TeamUsage) TableName() string {
	return "team_usages"
}