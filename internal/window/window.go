package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func RunUI() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(600, 400))

	tabs := getTabs()

	w.SetContent(tabs)
	w.ShowAndRun()
}

func getTabs() *container.AppTabs {
	tabs := container.NewAppTabs(
		container.NewTabItem("Invoice", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	return tabs
}
