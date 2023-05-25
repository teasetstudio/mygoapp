package invoice

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func getInvoiceContainer() *fyne.Container {
	label1 := widget.NewLabel("Label 1")
	value1 := widget.NewLabel("Value")
	label2 := widget.NewLabel("Label 2")
	value2 := widget.NewLabel("Something")
	grid := container.New(layout.NewFormLayout(), label1, value1, label2, value2)
	return grid
}
