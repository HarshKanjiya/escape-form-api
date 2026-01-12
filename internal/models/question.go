package models

type Question struct {
	ID            string           `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	FormID        string           `gorm:"type:uuid;index;column:form_id" json:"formId"`
	Title         string           `gorm:"column:title" json:"title"`
	Placeholder   string           `gorm:"default:'';column:placeholder" json:"placeholder"`
	Description   string           `gorm:"default:'';column:description" json:"description"`
	Required      bool             `gorm:"default:false;column:required" json:"required"`
	Type          QuestionType     `gorm:"column:type" json:"type"`
	Metadata      string           `gorm:"type:jsonb;default:'{}';column:metadata" json:"metadata"`
	PosX          int              `gorm:"column:pos_x" json:"posX"`
	PosY          int              `gorm:"column:pos_y" json:"posY"`
	SortOrder     *int             `gorm:"default:0;column:sort_order" json:"sortOrder"`
	Options       []QuestionOption `gorm:"foreignKey:QuestionID" json:"options"`
	OutgoingEdges []Edge           `gorm:"foreignKey:SourceNodeID" json:"outgoingEdges"`
	IncomingEdges []Edge           `gorm:"foreignKey:TargetNodeID" json:"incomingEdges"`
	Form          Form             `gorm:"references:ID" json:"form"`
}

func (Question) TableName() string {
	return "questions"
}