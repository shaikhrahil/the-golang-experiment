package team_user

import (
	"github.com/shaikhrahil/the-golang-experiment/rest"
	"github.com/shaikhrahil/the-golang-experiment/todo/lib/team"
	"github.com/shaikhrahil/the-golang-experiment/todo/lib/user"
)

type TeamUser struct {
	rest.Model
	TeamID uint      `json:"teamID"`
	Team   team.Team `json:"-"`
	UserID uint      `json:"userID"`
	User   user.User `json:"-"`
}
