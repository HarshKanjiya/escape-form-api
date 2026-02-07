package models

import "time"

type TeamSubscription struct {
	ID              string                   `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	TeamID          string                   `gorm:"type:uuid;column:teamId" json:"teamId"`
	PlanID          string                   `gorm:"type:uuid;column:planId" json:"planId"`
	Status          TeamSubscriptionStatus   `gorm:"default:ACTIVE;column:status" json:"status"`
	FormsAllowed    int                      `gorm:"default:1;column:formsAllowed" json:"formsAllowed"`
	ProjectsAllowed int                      `gorm:"default:1;column:projectsAllowed" json:"projectsAllowed"`
	Responses       int                      `gorm:"default:0;column:responses" json:"responses"`
	OverUsage       int                      `gorm:"default:0;column:overUsage" json:"overUsage"`
	CycleStart      time.Time                `gorm:"type:timestamptz(6);column:cycleStart" json:"cycleStart"`
	CycleEnd        time.Time                `gorm:"type:timestamptz(6);column:cycleEnd" json:"cycleEnd"`
	CancelCycleAt   *time.Time               `gorm:"type:timestamptz(6);column:cancelCycleAt" json:"cancelCycleAt"`
	CreatedAt       time.Time                `gorm:"type:timestamptz(6);default:now();column:createdAt" json:"createdAt"`
	UpdatedAt       time.Time                `gorm:"type:timestamptz(6);autoUpdateTime;column:updatedAt" json:"updatedAt"`
	Team            []Team                   `gorm:"foreignKey:TeamUsageID" json:"team"`
	Plan            Plan                     `gorm:"foreignKey:PlanID;references:ID;onDelete:CASCADE" json:"plan"`
}

func (TeamSubscription) TableName() string {
	return "team_subscriptions"
}