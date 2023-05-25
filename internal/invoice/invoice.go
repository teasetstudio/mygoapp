package invoice

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func RunUI() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
