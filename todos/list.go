package todos

import (
	"main/components"
	"net/http"

	"github.com/twharmon/godom"
)

type list struct {
	godom.Component
}

// List .
func List(ps godom.RouteParams) godom.Renderer {
	return &list{}
}

// Render .
func (r *list) Render(root *godom.Elem) {
	listContainer := godom.Create("div")
	root.AppendElem(listContainer)

	todoCh := make(chan []*Todo)
	btn := components.NewButton()
	godom.Mount(btn, root)
	btn.Text <- "Reload"
	btn.Handler <- func(e *godom.MouseEvent) {
		listContainer.Text("Loading...")
		r.getTodos(todoCh)
	}

	go r.getTodos(todoCh)

	go func() {
		for {
			select {
			case todos := <-todoCh:
				listContainer.Clear()
				for _, todo := range todos {
					listContainer.AppendElem(view(todo))
				}
			case <-r.Quit:
				return
			}
		}
	}()
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
