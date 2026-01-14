package models

import "time"

type ActivePassword struct {
	ID         string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	FormID     string    `gorm:"type:uuid;index;column:formId" json:"formId"`
	Name       string    `gorm:"type:varchar;column:name" json:"name"`
	Password   string    `gorm:"type:varchar;column:password" json:"password"`
	IsValid    bool      `gorm:"default:true;column:isValid" json:"isValid"`
	ExpireAt   time.Time `gorm:"type:timestamptz(6);column:expireAt" json:"expireAt"`
	CreatedAt  time.Time `gorm:"type:timestamptz(6);default:now();column:createdAt" json:"createdAt"`
	UsableUpto int       `gorm:"default:1;column:usableUpto" json:"usableUpto"`
	Form       Form      `gorm:"references:ID" json:"form"`
}

func (ActivePassword) TableName() string {
	return "active_passwords"
}
