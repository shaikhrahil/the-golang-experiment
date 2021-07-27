package todo

import "github.com/shaikhrahil/the-golang-experiment/rest"

type TodoSummary struct {
	rest.Base
	// UserID uint
	// TodoID uint
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
