package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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
	invoiceTab := getInvoiceContainer()
	invoiceDataTab := container.NewVScroll(getInvoiceDataContainer())
	tabs := container.NewAppTabs(
		container.NewTabItem("Data", invoiceDataTab),
		container.NewTabItem("Invoice", invoiceTab),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	return tabs
}
