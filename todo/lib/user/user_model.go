package user

import (
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
	todo "github.com/shaikhrahil/the-golang-experiment/todo/lib"
)

type User struct {
	accounts.User
	Todos []todo.Todo `gorm:"foreignKey:UserID"`
}
