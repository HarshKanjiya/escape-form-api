package models

type PlanFeature struct {
	ID        string  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	PlanID    string  `gorm:"type:uuid;index;column:planId" json:"planId"`
	FeatureID string  `gorm:"type:uuid;index;column:featureId" json:"featureId"`
	Plan      Plan    `gorm:"foreignKey:PlanID;references:ID;onDelete:CASCADE" json:"plan"`
	Feature   Feature `gorm:"foreignKey:FeatureID;references:ID;onDelete:CASCADE" json:"feature"`
}

func (PlanFeature) TableName() string {
	return "plan_features"
}
