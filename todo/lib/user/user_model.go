package user

import (
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
	"github.com/shaikhrahil/the-golang-experiment/todo/lib/team_user"
)

type User struct {
	accounts.User
	TeamUserIDs []team_user.TeamUser `gorm:"foreignKey:UserID"`
}
