package todos

import (
	"main/components"
	"net/http"

	"github.com/twharmon/godom"
)

type list struct {
	godom.BaseComponent
}

// List .
func List(ps godom.RouteParams) godom.Component {
	return &list{}
}

// Render .
func (r *list) Render() *godom.Elem {
	listContainer := godom.Create("div")

	btn := components.NewButton()
	go func() {
		btn.Text <- "Reload"
		btn.Handler <- func(e *godom.MouseEvent) {
			listContainer.Text("Loading...")
			r.getTodos(listContainer)
		}
	}()

	go r.getTodos(listContainer)
	go func() { <-r.Quit }()

	root := godom.Create("div")
	root.Append(btn, listContainer)
	return root
}

func (r *list) getTodos(e *godom.Elem) {
	var todos []*Todo
	err := r.HTTP(http.MethodGet, "https://jsonplaceholder.typicode.com/todos", nil).FromJSON(&todos, func(res *http.Response) {
		switch res.StatusCode {
		case http.StatusOK:
			e.Clear()
			for _, todo := range todos {
				e.Append(view(todo))
			}
		default:
			godom.Log(res.Status)
		}
	})
	if err != nil {
		godom.Log(err.Error())
	}
}
