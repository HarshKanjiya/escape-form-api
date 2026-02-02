package models

import "gorm.io/datatypes"

type Question struct {
	ID            string           `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	FormID        string           `gorm:"type:uuid;index;column:formId" json:"formId"`
	Title         string           `gorm:"column:title" json:"title"`
	Placeholder   string           `gorm:"default:'';column:placeholder" json:"placeholder"`
	Description   string           `gorm:"default:'';column:description" json:"description"`
	Required      bool             `gorm:"default:false;column:required" json:"required"`
	Type          QuestionType     `gorm:"column:type" json:"type"`
	Metadata      datatypes.JSON   `gorm:"type:jsonb;default:'{}';column:metadata" json:"metadata"`
	PosX          int              `gorm:"column:posX" json:"posX"`
	PosY          int              `gorm:"column:posY" json:"posY"`
	SortOrder     *int             `gorm:"default:0;column:sortOrder" json:"sortOrder"`
	Options       []QuestionOption `gorm:"foreignKey:QuestionID" json:"options"`
	OutgoingEdges []Edge           `gorm:"foreignKey:SourceNodeID" json:"outgoingEdges"`
	IncomingEdges []Edge           `gorm:"foreignKey:TargetNodeID" json:"incomingEdges"`
	Form          Form             `gorm:"foreignKey:FormID;references:ID;onDelete:CASCADE" json:"form"`
}

func (Question) TableName() string {
	return "questions"
}
