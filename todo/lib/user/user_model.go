package user

import (
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
	"github.com/shaikhrahil/the-golang-experiment/rest"
	todo "github.com/shaikhrahil/the-golang-experiment/todo/lib"
)

type User struct {
	accounts.User
	Todos []todo.Todo `gorm:"foreignKey:UserID"`
}

type UserTodo struct {
	rest.Base
	UserID uint `gorm:"primaryKey"`
	TodoID uint `gorm:"primaryKey"`
}
