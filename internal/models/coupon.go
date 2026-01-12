package models

type Coupon struct {
	ID           string             `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	Type         CouponType         `gorm:"default:'GENERAL';column:type" json:"type"`
	DiscountType CouponDiscountType `gorm:"default:'FLAT';column:discount_type" json:"discountType"`
	PlanID       *string            `gorm:"type:uuid;column:plan_id" json:"planId"`
	Amount       *float64           `gorm:"column:amount" json:"amount"`
	UseLeft      int                `gorm:"default:0;column:use_left" json:"useLeft"`
	Plan         *Plan              `gorm:"references:ID" json:"plan"`
}

func (Coupon) TableName() string {
	return "coupons"
}