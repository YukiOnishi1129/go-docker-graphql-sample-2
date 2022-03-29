package entity

type Todo struct {
	BaseModel
	Title   string `gorm:"size:255" json:"title,omitempty"`
	Comment string `gorm:"type:text" json:"comment,omitempty"`
}