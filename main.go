package main

import (
	"main/app"

	"github.com/twharmon/godom"
)

func main() {
	app := app.New()
	godom.Root("#root").Append(app)
	<-app.Quit
}
