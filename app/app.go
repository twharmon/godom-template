package app

import (
	"fmt"
	"main/hello"
	"main/home"
	"main/todos"

	"github.com/twharmon/godom"
)

// App .
type App struct {
	godom.BaseComponent
}

// New .
func New() *App {
	return &App{}
}

// Render .
func (a *App) Render() *godom.Elem {
	router := godom.NewRouter()

	router.On("/", home.New)
	router.On("/hello/{name}", hello.New)
	router.On("/todos", todos.List)

	outlet := godom.Create("div")
	router.Mount(outlet)

	return godom.Create("div").Append(nav(), outlet)
}

func nav() *godom.Elem {
	nav := godom.Create("ul")
	nav.Append(
		navItem("Home", "/"),
		navItem("Hello", "/hello/Jimmy"),
		navItem("Todos", "/todos"),
	)
	return nav
}

func navItem(text string, href string) *godom.Elem {
	a := godom.Create("a").Text(text).Attr("href", fmt.Sprintf("/#%s", href))
	li := godom.Create("li")
	li.Append(a)
	return li
}
