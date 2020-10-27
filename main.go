package main

import (
	"main/app"

	"github.com/twharmon/godom"
)

func main() {
	app := app.New()
	godom.Mount(app, godom.Root("#root"))
	<-app.Quit
}
