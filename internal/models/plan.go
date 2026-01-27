package models

import "time"

type Plan struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Description *string   `gorm:"column:description" json:"description"`
	Price       float64   `gorm:"column:price" json:"price"`
	Currency    string    `gorm:"type:varchar(3);default:'INR';column:currency" json:"currency"`
	Valid       bool      `gorm:"default:true;column:valid" json:"valid"`
	CreatedAt   time.Time `gorm:"type:timestamptz(6);default:now();column:createdAt" json:"createdAt"`
	UpdatedAt   *time.Time `gorm:"type:timestamptz(6);autoUpdateTime;column:updatedAt" json:"updatedAt"`
	Coupons     []Coupon  `gorm:"foreignKey:PlanID" json:"coupons"`
	Features    []Feature `gorm:"foreignKey:PlanID" json:"features"`
	Teams       []Team    `gorm:"foreignKey:PlanID" json:"teams"`
}

func (Plan) TableName() string {
	return "plans"
}