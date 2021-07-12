package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `validate:"required,min=3,max=32" partial_validate:"omitempty,min=3,max=32"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" validate:"required,email,min=3,max=32" partial_validate:"omitempty,email,min=3,max=32"`
	Username string `gorm:"uniqueIndex;type:varchar(255)" validate:"required,min=3,max=32" partial_validate:"omitempty,min=3,max=32"`
}
