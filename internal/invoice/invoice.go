package invoice

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func GetInvoiceContainer() *fyne.Container {
	renameInvoice := RenameInvoiceContainer()

	grid := container.NewVBox(renameInvoice)
	return grid
}
