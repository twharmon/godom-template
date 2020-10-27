package home

import (
	"fmt"
	"main/data"

	"github.com/twharmon/godom"
)

// Home .
type Home struct {
	godom.Component
}

// New .
func New(ps godom.RouteParams) godom.Renderer {
	return &Home{}
}

// Render .
func (h *Home) Render(root *godom.Elem) {
	p := godom.Create("p")
	root.AppendElem(p)

	state := data.Store.Subscribe()

	go func() {
		for {
			select {
			case s := <-state:
				p.Text(fmt.Sprintf("%s's home", s.User.Name))
			case <-h.Quit:
				data.Store.Unsubscribe(state)
				return
			}
		}
	}()
}
