package todo

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deleteddAt" gorm:"index"`

	Title   string `json:"title" gorm:"type:varchar(255)" validate:"omitempty,min=3,max=32"`
	Content string `json:"content" gorm:"type:varchar(255)" validate:"omitempty,min=3,max=32"`
	Done    bool   `json:"done" gorm:"type:boolean,default:false"`
}
