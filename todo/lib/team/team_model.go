package team

import (
	"github.com/shaikhrahil/the-golang-experiment/rest"
	"github.com/shaikhrahil/the-golang-experiment/todo/lib/user"
)

type Team struct {
	rest.Model
	Name        string      `json:"name" gorm:"type:varchar(255)" validate:"omitempty,max=32"`
	Description string      `json:"description" gorm:"type:varchar(255)" validate:"omitempty"`
	Users       []user.User `json:"users" gorm:"many2many:team_users"`
}
