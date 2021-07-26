package todo

import (
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
	"github.com/shaikhrahil/the-golang-experiment/rest"
)

type User struct {
	accounts.User
	Todos []Todo `gorm:"many2many:user_todos;"`
}

type UserTodo struct {
	rest.Base
	UserID int `gorm:"primaryKey"`
	TodoID int `gorm:"primaryKey"`
}
