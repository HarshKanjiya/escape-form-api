package models

type Coupon struct {
	ID           string             `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	Type         CouponType         `gorm:"default:'GENERAL';column:type" json:"type"`
	DiscountType CouponDiscountType `gorm:"default:'FLAT';column:discountType" json:"discountType"`
	PlanID       *string            `gorm:"type:uuid;column:planId" json:"planId"`
	Amount       *float64           `gorm:"column:amount" json:"amount"`
	UseLeft      int                `gorm:"default:0;column:useLeft" json:"useLeft"`
	Plan         *Plan              `gorm:"foreignKey:PlanID;references:ID" json:"plan"`
}

func (Coupon) TableName() string {
	return "coupons"
}