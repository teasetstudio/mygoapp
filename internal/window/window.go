package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

var mainWindow fyne.Window

func RunUI() {
	a := app.New()
	mainWindow = a.NewWindow("Hello World")
	mainWindow.Resize(fyne.NewSize(600, 400))

	tabs := getTabs()

	mainWindow.SetContent(tabs)
	mainWindow.ShowAndRun()
}

func getTabs() *container.AppTabs {
	invoiceTab := getInvoiceContainer()
	invoiceDataTab := container.NewVScroll(getInvoiceDataContainer())
	tabs := container.NewAppTabs(
		container.NewTabItem("Invoice", invoiceTab),
		container.NewTabItem("Data", invoiceDataTab),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	return tabs
}
