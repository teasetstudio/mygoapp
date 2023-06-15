package invoice

import (
	"fmt"
	"log"
	"mygoapp/internal/config"
	"mygoapp/internal/scrap"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

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
	date, year, month, monthStr := getInvoiceDate()
	DefaultFolderSetup()

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
		errr := os.MkdirAll(filepath.Dir(downloadFilePath), 0755)
		if errr != nil {
			fmt.Println("Error creating directory:", errr)
		}

		scrap.GetInvoice(invoiceDate, downloadFilePath)

		// err := os.Rename(downloadedFilePath, newFileInput.Text)
		// if os.IsNotExist(err) {
		// 	log.Println("Error: File not found")
		// }
		// if os.IsExist(err) {
		// 	log.Println("Error: File already exists")
		// }
		// if err != nil && !os.IsNotExist(err) && !os.IsExist(err) {
		// 	log.Println(err)
		// }
		// if err == nil {
		// 	log.Println("Success: File renamed")
		// }
		// email.SendEmailWithFile(newFileInput.Text)
	})
	btnWrapper := container.New(layout.NewCenterLayout(), btn)

	grid := container.NewVBox(form, btnWrapper)
	return grid
}

func getInvoiceDate() (string, string, string, string) {
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

	return strconv.Itoa(lastDayOfMonth) + "-" + mon + "-" + strconv.Itoa(year), strconv.Itoa(int(year)), stringMonth, month.String()
}

func daysIn(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
