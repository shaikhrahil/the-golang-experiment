package team_user

import "github.com/shaikhrahil/the-golang-experiment/rest"

type TeamUser struct {
	rest.Model
	TeamID uint `json:"teamID" gorm:"primaryKey"`
	UserID uint `json:"userID" gorm:"primaryKey"`
}
