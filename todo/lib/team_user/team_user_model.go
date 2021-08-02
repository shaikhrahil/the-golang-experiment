package team_user

import (
	"github.com/shaikhrahil/the-golang-experiment/rest"
	"github.com/shaikhrahil/the-golang-experiment/todo/lib/team"
	"github.com/shaikhrahil/the-golang-experiment/todo/lib/user"
)

type TeamUser struct {
	rest.Model
	TeamID uint      `json:"teamID" gorm:"unique"`
	Team   team.Team `json:"-"`
	UserID uint      `json:"userID" gorm:"unique"`
	User   user.User `json:"-"`
}
