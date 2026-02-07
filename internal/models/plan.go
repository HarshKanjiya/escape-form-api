package models

import "time"

type Plan struct {
	ID               string              `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	Name             string              `gorm:"column:name" json:"name"`
	Description      *string             `gorm:"column:description" json:"description"`
	RefID            *string             `gorm:"type:varchar;column:refId" json:"refId"`
	PriceMonthly     float64             `gorm:"column:priceMonthly" json:"priceMonthly"`
	PriceYearly      float64             `gorm:"column:priceYearly" json:"priceYearly"`
	YearlyDiscount   *float64            `gorm:"column:yearlyDiscount" json:"yearlyDiscount"`
	MaxForms         *int                `gorm:"column:maxForms" json:"maxForms"`
	MaxProjects      *int                `gorm:"column:maxProjects" json:"maxProjects"`
	MaxSubmission    *int                `gorm:"column:maxSubmission" json:"maxSubmission"`
	MaxFileStorage   *float64            `gorm:"column:maxFileStorage" json:"maxFileStorage"`
	MaxTeamMembers   *int                `gorm:"column:maxTeamMembers" json:"maxTeamMembers"`
	MaxTokens        *int                `gorm:"column:maxTokens" json:"maxTokens"`
	Valid            bool                `gorm:"default:true;column:valid" json:"valid"`
	CreatedAt        time.Time           `gorm:"type:timestamptz(6);default:now();column:createdAt" json:"createdAt"`
	UpdatedAt        *time.Time          `gorm:"type:timestamptz(6);autoUpdateTime;column:updatedAt" json:"updatedAt"`
	Teams            []Team              `gorm:"foreignKey:PlanID" json:"teams"`
	PlanFeatures     []PlanFeature       `gorm:"foreignKey:PlanID" json:"planFeatures"`
	TeamSubscriptions []TeamSubscription `gorm:"foreignKey:PlanID" json:"teamSubscriptions"`
}

func (Plan) TableName() string {
	return "plans"
}
