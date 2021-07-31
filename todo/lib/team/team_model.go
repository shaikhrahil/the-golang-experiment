package team

import (
	"github.com/shaikhrahil/the-golang-experiment/rest"
	"github.com/shaikhrahil/the-golang-experiment/todo/lib/team_user"
)

type Team struct {
	rest.Model
	Name        string               `json:"name" gorm:"type:varchar(255)" validate:"omitempty,max=32"`
	Description string               `json:"description" gorm:"type:varchar(255)" validate:"omitempty"`
	TeamUserIDs []team_user.TeamUser `gorm:"foreignKey:TeamID"`
}
