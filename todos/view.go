package todos

import (
	"github.com/twharmon/godom"
)

// Todo .
type Todo struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func view(todo *Todo) *godom.Elem {
	div := godom.Create("div")
	div.Text(todo.Title)
	div.Classes("todo")
	if todo.Completed {
		div.AddClass("completed")
	}
	return div
}
