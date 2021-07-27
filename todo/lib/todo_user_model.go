package todo

import (
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
	"github.com/shaikhrahil/the-golang-experiment/rest"
)

type User struct {
	accounts.User
}

type UserTodo struct {
	rest.Base
	UserID uint `gorm:"primaryKey"`
	TodoID uint `gorm:"primaryKey"`
}
