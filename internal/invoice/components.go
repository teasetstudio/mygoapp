package invoice

import (
	"log"
	"mygoapp/internal/config"
	"mygoapp/internal/scrap"
	"os"
	"runtime"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// func FormContainer() *fyne.Container {
// 	label1 := widget.NewLabel("Label 1")
// 	input1 := widget.NewEntry()
// 	input1.SetPlaceHolder("Enter text...")

// 	label2 := widget.NewLabel("Label 2")
// 	input2 := widget.NewEntry()
// 	input2.SetPlaceHolder("Enter text...")

// 	form := container.New(layout.NewFormLayout(), label1, input1, label2, input2)

// 	return form
// }

func DefaultFolderSetup() {
	if runtime.GOOS == "windows" {
		if err := os.Mkdir(config.AppDir, os.ModePerm); err != nil {
			if os.IsExist(err) {
				// check that the existing path is a directory
				info, err := os.Stat(config.AppDir)
				if err != nil {
					log.Fatal(err)
				}
				if !info.IsDir() {
					log.Fatal("path exists but is not a directory")
				}
			} else {
				log.Fatal(err)
			}
		}
	} else {
		log.Fatal("OS not supported")
	}
}

func RenameInvoiceContainer() *fyne.Container {
	date := getInvoiceDate()
	DefaultFolderSetup()
	// download := config.DownloadDir + "\\faktura-vat-" + date + ".pdf"
	newFile := config.DownloadDir + "\\May Godel Invoice I.Tichkevitch.pdf"

	invoiceDateLabel := widget.NewLabel("Invoice date")
	invoiceDateInput := widget.NewEntry()
	invoiceDateInput.SetText(date)

	downloadFileLabel := widget.NewLabel("Download path")
	downloadFileInput := widget.NewEntry()
	downloadFileInput.SetText(config.DownloadDir)

	newFileLabel := widget.NewLabel("New file path")
	newFileInput := widget.NewEntry()
	newFileInput.SetText(newFile)

	form := container.New(layout.NewFormLayout(), invoiceDateLabel, invoiceDateInput, downloadFileLabel, downloadFileInput, newFileLabel, newFileInput)

	btn := widget.NewButton("Save", func() {
		invoiceDate := invoiceDateInput.Text
		downloadedFilePath := scrap.Start(invoiceDate)

		println(downloadedFilePath)
		println(newFileInput.Text)
		err := os.Rename(downloadedFilePath, newFileInput.Text)
		if os.IsNotExist(err) {
			log.Println("Error: File not found")
		}
		if os.IsExist(err) {
			log.Println("Error: File already exists")
		}
		if err != nil && !os.IsNotExist(err) && !os.IsExist(err) {
			log.Println(err)
		}
		if err == nil {
			log.Println("Success: File renamed")
		}
	})
	btnWrapper := container.New(layout.NewCenterLayout(), btn)

	grid := container.NewVBox(form, btnWrapper)
	return grid
}

func getInvoiceDate() string {
	year, month, _ := time.Now().Date()

	lastDayOfMonth := daysIn(month, year)
	stringMonth := strconv.Itoa(int(month))

	mon := func() string {
		if month > 9 {
			return stringMonth
		} else {
			return "0" + stringMonth
		}
	}()

	return strconv.Itoa(lastDayOfMonth) + "-" + mon + "-" + strconv.Itoa(year)
}

func daysIn(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
