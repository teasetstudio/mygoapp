package window

import (
	"fmt"
	"mygoapp/internal/config"
	"mygoapp/internal/email"
	"mygoapp/internal/pdf"
	"mygoapp/internal/scrap"
	"os"
	"os/exec"
	"path/filepath"

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

	sendEmailLabel := widget.NewLabel("")
	isSendEmail := true
	sendEmailCheckbox := widget.NewCheck("Do want to send email?", func(value bool) {
		isSendEmail = value
	})
	sendEmailCheckbox.Checked = isSendEmail
	sendEmailCheckbox.Refresh()

	form := container.New(layout.NewFormLayout(), invoiceDateLabel, invoiceDateInput, downloadFileLabel, downloadFileInput, sendEmailLabel, sendEmailCheckbox)

	btn := widget.NewButton("Save", func() {
		defer recoverDialog()

		invoiceDate := invoiceDateInput.Text
		downloadFilePath := downloadFileInput.Text

		errF := os.MkdirAll(filepath.Dir(downloadFilePath), 0755)
		if errF != nil {
			fmt.Println("Error creating directory:", errF)
		}

		pdf.GetPdfInvoice(invoiceDate, downloadFilePath, &config.InvoiceData)
		// Unnesessary anymore
		if false {
			scrap.DownloadInvoice(invoiceDate, downloadFilePath, &config.InvoiceData)
		}
		if isSendEmail {
			email.SendFile(downloadFilePath, &config.InvoiceData.User)
		}

		err := exec.Command("explorer", downloadFilePath).Start()
		if err != nil {
			panic(err)
		}
		showDialog("Success", "File created")
	})
	btnWrapper := container.New(layout.NewCenterLayout(), btn)

	grid := container.NewVBox(form, btnWrapper)
	return grid
}
