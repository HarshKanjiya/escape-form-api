package models

import "time"

type AddOn struct {
	ID          string      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	Key         string      `gorm:"column:key" json:"key"`
	Name        string      `gorm:"column:name" json:"name"`
	Description *string     `gorm:"column:description" json:"description"`
	RefID       *string     `gorm:"type:varchar;column:refId" json:"refId"`
	Price       float64     `gorm:"column:price" json:"price"`
	Valid       bool        `gorm:"default:true;column:valid" json:"valid"`
	CreatedAt   time.Time   `gorm:"type:timestamptz(6);default:now();column:createdAt" json:"createdAt"`
	UpdatedAt   *time.Time  `gorm:"type:timestamptz(6);autoUpdateTime;column:updatedAt" json:"updatedAt"`
	TeamAddons  []TeamAddon `gorm:"foreignKey:AddOnID" json:"teamAddons"`
}

func (AddOn) TableName() string {
	return "add_ons"
}
