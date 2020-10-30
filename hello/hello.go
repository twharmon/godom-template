package hello

import (
	"main/data"

	"github.com/twharmon/godom"
)

// Hello .
type Hello struct {
	godom.BaseComponent
	name string
}

// New .
func New(ps godom.RouteParams) godom.Component {
	name := ps.Get("name")
	data.Store.SetUser(data.User{ID: 5, Name: name})
	return &Hello{
		name: name,
	}
}

// Render .
func (r *Hello) Render() *godom.Elem {
	p := godom.Create("p").Text(r.name)
	go func() { <-r.Quit }()
	return p
}
