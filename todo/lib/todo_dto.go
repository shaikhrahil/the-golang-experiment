package todo

import "github.com/shaikhrahil/the-golang-experiment/rest"

type TodoSummary struct {
	rest.Model
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
