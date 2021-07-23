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

	Title   string `json:"title" validate:"omitempty,min=3,max=32" gorm:"type:varchar(255)"`
	Content string `json:"content" validate:"omitempty,min=3,max=32" gorm:"type:varchar(255)"`
}
