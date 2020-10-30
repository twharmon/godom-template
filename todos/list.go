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
	root := godom.Create("div")
	listContainer := godom.Create("div")

	todoCh := make(chan []*Todo)
	btn := components.NewButton()
	go func() {
		btn.Text <- "Reload"
		btn.Handler <- func(e *godom.MouseEvent) {
			listContainer.Text("Loading...")
			r.getTodos(todoCh)
		}
	}()

	root.Append(btn, listContainer)

	go r.getTodos(todoCh)

	go func() {
		for {
			select {
			case todos := <-todoCh:
				listContainer.Clear()
				for _, todo := range todos {
					listContainer.Append(view(todo))
				}
			case <-r.Quit:
				return
			}
		}
	}()

	return root
}

func (r *list) getTodos(todoCh chan []*Todo) {
	var todos []*Todo
	err := r.HTTP(http.MethodGet, "https://jsonplaceholder.typicode.com/todos", nil).FromJSON(&todos, func(res *http.Response) {
		switch res.StatusCode {
		case http.StatusOK:
			todoCh <- todos
		default:
			godom.Log(res.Status)
		}
	})
	if err != nil {
		godom.Log(err.Error())
	}
}
