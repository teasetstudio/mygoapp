package window

import (
	"mygoapp/internal/config"
	"mygoapp/internal/email"
	"mygoapp/internal/scrap"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func getInvoiceContainer() *fyne.Container {
	invoiceContainer := invoiceContainer()

	grid := container.NewVBox(invoiceContainer)
	return grid
}

func invoiceContainer() *fyne.Container {
	date, year, month, monthStr := getInvoiceDate()

	monthFolderName := month + " " + monthStr
	newFileFolder := config.BusinessDir + "\\" + year + "\\" + monthFolderName + "\\" + "invoices"
	newFile := newFileFolder + "\\" + monthStr + " Godel Invoice I.Tichkevitch.pdf"

	invoiceDateLabel := widget.NewLabel("Invoice date")
	invoiceDateInput := widget.NewEntry()
	invoiceDateInput.SetText(date)

	downloadFileLabel := widget.NewLabel("Download path")
	downloadFileInput := widget.NewEntry()
	downloadFileInput.SetText(newFile)

	form := container.New(layout.NewFormLayout(), invoiceDateLabel, invoiceDateInput, downloadFileLabel, downloadFileInput)

	btn := widget.NewButton("Save", func() {
		invoiceDate := invoiceDateInput.Text
		downloadFilePath := downloadFileInput.Text

		scrap.DownloadInvoice(invoiceDate, downloadFilePath, &config.InvoiceData)
		email.SendFile(downloadFilePath)
	})
	btnWrapper := container.New(layout.NewCenterLayout(), btn)

	grid := container.NewVBox(form, btnWrapper)
	return grid
}
