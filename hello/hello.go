package hello

import (
	"main/data"

	"github.com/twharmon/godom"
)

// Hello .
type Hello struct {
	godom.Component
	name string
}

// New .
func New(ps godom.RouteParams) godom.Renderer {
	name := ps.Get("name")
	data.Store.SetUser(data.User{ID: 5, Name: name})
	return &Hello{
		name: name,
	}
}

// Render .
func (r *Hello) Render(root *godom.Elem) {
	p := godom.Create("p").Text(r.name)
	root.AppendElem(p)
	go func() { <-r.Quit }()
}
