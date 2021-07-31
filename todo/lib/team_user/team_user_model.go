package team_user

import "github.com/shaikhrahil/the-golang-experiment/rest"

type TeamUser struct {
	rest.MapModel
	TeamID uint `json:"teamID"`
	UserID uint `json:"userID"`
}
