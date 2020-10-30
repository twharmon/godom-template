package home

import (
	"fmt"
	"main/data"

	"github.com/twharmon/godom"
)

// Home .
type Home struct {
	godom.BaseComponent
}

// New .
func New(ps godom.RouteParams) godom.Component {
	return &Home{}
}

// Render .
func (h *Home) Render() *godom.Elem {
	p := godom.Create("p")

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

	return p
}
