package models

type Edge struct {
	ID           string   `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	FormID       string   `gorm:"type:uuid;index;column:formId" json:"formId"`
	SourceNodeID string   `gorm:"type:uuid;column:sourceNodeId" json:"sourceNodeId"`
	TargetNodeID string   `gorm:"type:uuid;column:targetNodeId" json:"targetNodeId"`
	Condition    *string  `gorm:"type:jsonb;default:'{}';column:condition" json:"condition"`
	Form         Form     `gorm:"references:ID" json:"form"`
	SourceNode   Question `gorm:"references:ID" json:"sourceNode"`
	TargetNode   Question `gorm:"references:ID" json:"targetNode"`
}

func (Edge) TableName() string {
	return "edges"
}
