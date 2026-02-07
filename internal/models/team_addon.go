package models

type TeamAddon struct {
	ID       string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	TeamID   string `gorm:"type:uuid;index;column:teamId" json:"teamId"`
	AddOnID  string `gorm:"type:uuid;index;column:addOnId" json:"addOnId"`
	Quantity int    `gorm:"default:1;column:quantity" json:"quantity"`
	Team     Team   `gorm:"foreignKey:TeamID;references:ID;onDelete:CASCADE" json:"team"`
	AddOn    AddOn  `gorm:"foreignKey:AddOnID;references:ID;onDelete:CASCADE" json:"addOn"`
}

func (TeamAddon) TableName() string {
	return "team_addons"
}
