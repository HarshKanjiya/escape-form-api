package models

type QuestionOption struct {
	ID         string   `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id" json:"id"`
	QuestionID string   `gorm:"type:uuid;index;column:questionId" json:"questionId"`
	Label      string   `gorm:"column:label" json:"label"`
	Value      string   `gorm:"column:value" json:"value"`
	SortOrder  int      `gorm:"default:0;column:sortOrder" json:"sortOrder"`
	Question   Question `gorm:"foreignKey:QuestionID;references:ID;onDelete:CASCADE" json:"question"`
}

func (QuestionOption) TableName() string {
	return "questions_options"
}
