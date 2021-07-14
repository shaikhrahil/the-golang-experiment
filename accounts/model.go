package main

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deleteddAt" gorm:"index"`

	FirstName string `json:"firstName" validate:"required,min=3,max=32" partial_validate:"omitempty,min=3,max=32"`
	LastName  string `json:"lastName" validate:"omitempty,min=3,max=32"`
	Email     string `json:"email" gorm:"uniqueIndex;type:varchar(255)" validate:"required,email,min=3,max=32" partial_validate:"omitempty,email,min=3,max=32"`
	UserName  string `json:"userName" gorm:"uniqueIndex;type:varchar(255)" validate:"required,min=3,max=32" partial_validate:"omitempty,min=3,max=32"`
}
