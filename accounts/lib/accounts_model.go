package accounts

import (
	"github.com/shaikhrahil/the-golang-experiment/rest"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	rest.Base

	FirstName string `json:"firstName" gorm:"type:varchar(255)" validate:"required,min=3,max=32" partial_validate:"omitempty,min=3,max=32"`
	LastName  string `json:"lastName" gorm:"type:varchar(255)" validate:"omitempty,min=3,max=32"`
	Email     string `json:"email" gorm:"uniqueIndex;type:varchar(255)" validate:"required,email,min=3,max=32" partial_validate:"omitempty,email,min=3,max=32"`
	UserName  string `json:"userName" gorm:"uniqueIndex;type:varchar(255)" validate:"required,min=3,max=32" partial_validate:"omitempty,min=3,max=32"`
	Password  string `json:"-" gorm:"uniqueIndex;type:varchar(255)" validate:"required,min=10,max=32" partial_validate:"omitempty,min=10,max=32"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)

	if err != nil {
		panic(err)
	}

	u.Password = string(hash)

	return
}
