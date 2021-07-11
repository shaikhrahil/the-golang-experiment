package main

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string     `json:"name" validate:"required,min=3,max=32"`
	Email     string     `json:"email" gorm:"uniqueIndex,type:varchar(255)" validate:"required,email,min=3,max=32"`
	Username  string     `json:"username" gorm:"uniqueIndex,type:varchar(255)" validate:"required,min=3,max=32"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
