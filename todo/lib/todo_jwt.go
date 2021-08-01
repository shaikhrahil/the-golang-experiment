package todo

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
	"github.com/shaikhrahil/the-golang-experiment/todo/lib/team_user"
)

func NewClaims(teamUsersRepo team_user.Repository, logger *log.Logger) func(payload accounts.User) jwt.MapClaims {
	return func(payload accounts.User) jwt.MapClaims {
		var teamUsers []team_user.TeamUser
		if err := teamUsersRepo.GetTeamsOfUser(&teamUsers, payload.ID).Error; err != nil {
			logger.Fatalln(err.Error())
		}
		var teams []uint
		for _, tU := range teamUsers {
			teams = append(teams, tU.TeamID)
		}
		p := jwt.MapClaims{
			"UserID": payload.ID,
			"Teams":  teams,
		}
		return p
	}
}
