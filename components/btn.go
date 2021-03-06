package components

import (
	"github.com/twharmon/godom"
)

// Button .
type Button struct {
	godom.BaseComponent
	Text     chan string
	Handler  chan func(*godom.MouseEvent)
	Disabled chan bool
}

// NewButton .
func NewButton() *Button {
	return &Button{
		Text:     make(chan string),
		Disabled: make(chan bool),
		Handler:  make(chan func(*godom.MouseEvent)),
	}
}

// Render .
func (b *Button) Render() *godom.Elem {
	btn := godom.Create("button")

	go func() {
		for {
			select {
			case text := <-b.Text:
				btn.Text(text)
			case disabled := <-b.Disabled:
				btn.Attr("disabled", disabled)
			case handler := <-b.Handler:
				btn.OnClick(handler)
			case <-b.Quit:
				return
			}
		}
	}()

	return btn
}
