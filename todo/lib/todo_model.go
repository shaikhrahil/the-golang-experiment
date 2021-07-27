package todo

import (
	"github.com/shaikhrahil/the-golang-experiment/rest"
	"gorm.io/gorm"
)

type Todo struct {
	rest.Base
	Title      string   `json:"title" gorm:"type:varchar(255)" validate:"omitempty,min=3,max=32"`
	Content    string   `json:"content" gorm:"type:varchar(255)" validate:"omitempty,min=3,max=32"`
	Done       bool     `json:"done" gorm:"default:false"`
	UserTodoID uint     `json:"-"`
	UserTodo   UserTodo `json:"-" gorm:"foreignKey:UserTodoID"`
}

func (u *Todo) AfterCreate(tx *gorm.DB) (err error) {
	u.UserTodo.TodoID = u.ID
	return
}
