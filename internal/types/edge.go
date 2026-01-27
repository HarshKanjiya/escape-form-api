package types

type EdgeDto struct {
	ID           string  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	FormID       string  `gorm:"type:uuid;index;column:formId" json:"formId"`
	SourceNodeID string  `gorm:"type:uuid;column:sourceNodeId" json:"sourceNodeId"`
	TargetNodeID string  `gorm:"type:uuid;column:targetNodeId" json:"targetNodeId"`
	Condition    *string `gorm:"type:jsonb;default:'{}';column:condition" json:"condition"`
}
