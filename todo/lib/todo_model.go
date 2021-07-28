package todo

import (
	"github.com/shaikhrahil/the-golang-experiment/rest"
)

type Todo struct {
	rest.Base
	Title   string `json:"title" gorm:"type:varchar(255)" validate:"omitempty,min=3,max=32"`
	Content string `json:"content" gorm:"type:varchar(255)" validate:"omitempty,min=3,max=32"`
	Done    bool   `json:"done" gorm:"default:false"`
	UserID  uint   `json:"userID"`
}
