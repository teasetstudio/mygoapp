package scrap

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"mygoapp/internal/config"

	"github.com/mxschmitt/playwright-go"
)

func DownloadInvoice(invoiceDate string, downloadFilePath string, invoiceConfig *config.InvoiceDataType) {
	err := os.MkdirAll(filepath.Dir(downloadFilePath), 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
	}

	pw, err := playwright.Run(&playwright.RunOptions{Browsers: []string{"chromium"}})
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(true)})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	page, err := browser.NewPage(playwright.BrowserNewContextOptions{AcceptDownloads: playwright.Bool(true)})
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	fakturowoLogin(page, invoiceConfig)
	if _, err = page.Goto("https://www.fakturowo.pl/wystaw", playwright.PageGotoOptions{Timeout: playwright.Float(100000)}); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	// p:contains("example")
	// nabywcaInput, err := page.QuerySelector("#nazwa_nabywca")
	// if err != nil {
	// 	log.Fatalf("could not get Nazwa: %v", err)
	// }
	// nabywcaInput.Fill(config.Nabywca.Nazwa)
	// Sprzedawca
	page.Fill("#nazwa_sprzedawca", invoiceConfig.Sprzedawca.Nazwa)
	page.Fill("#nip_sprzedawca", invoiceConfig.Sprzedawca.Nip)
	page.Fill("#ulica_sprzedawca", invoiceConfig.Sprzedawca.Ulica)
	page.Fill("#budynek_sprzedawca", invoiceConfig.Sprzedawca.Nr_budynku)
	if invoiceConfig.Sprzedawca.Lokalu != "" {
		page.Fill("#lokal_sprzedawca", invoiceConfig.Sprzedawca.Lokalu)
	}
	page.Fill("#miasto_sprzedawca", invoiceConfig.Sprzedawca.Miasto)
	page.Fill("#kod_sprzedawca", invoiceConfig.Sprzedawca.Kod)
	// Nabywca
	page.Fill("#nazwa_nabywca", invoiceConfig.Nabywca.Nazwa)
	page.Fill("#nip_nabywca", invoiceConfig.Nabywca.Nip)
	page.Fill("#ulica_nabywca", invoiceConfig.Nabywca.Ulica)
	page.Fill("#budynek_nabywca", invoiceConfig.Nabywca.Nr_budynku)
	if invoiceConfig.Nabywca.Lokalu != "" {
		page.Fill("#lokal_nabywca", invoiceConfig.Nabywca.Lokalu)
	}
	page.Fill("#miasto_nabywca", invoiceConfig.Nabywca.Miasto)
	page.Fill("#kod_nabywca", invoiceConfig.Nabywca.Kod)
	// Towar
	page.Fill("#nazwa_0", invoiceConfig.Towar.Nazwa)
	cena := strconv.FormatFloat(invoiceConfig.Towar.Cena, 'f', 2, 64)
	page.Fill("#cena_netto_0", cena)

	podpisCheckbox1, err := page.QuerySelector("input[name=\"sprzedawca[pokaz_podpis]\"]")
	if err != nil {
		log.Fatalf("could not get podpisCheckbox1: %v", err)
	}
	if is, _ := podpisCheckbox1.IsChecked(); is {
		podpisCheckbox1.Click()
	}

	podpisCheckbox2, _ := page.QuerySelector("input[name=\"nabywca[pokaz_podpis]\"]")
	if is, _ := podpisCheckbox2.IsChecked(); is {
		podpisCheckbox2.Click()
	}

	terminCheckbox, _ := page.QuerySelector("input[name=\"pokaz_termin\"]")
	if is, _ := terminCheckbox.IsChecked(); is {
		terminCheckbox.Click()
	}

	uwagiCheckbox, _ := page.QuerySelector("input[name=\"pokaz_uwagi\"]")
	if is, _ := uwagiCheckbox.IsChecked(); is {
		uwagiCheckbox.Click()
	}

	pokazCheckbox, _ := page.QuerySelector("input[name=pokaz_miejsce]")
	if is, _ := pokazCheckbox.IsChecked(); is {
		pokazCheckbox.Click()
	}

	page.Fill("input[name=numer]", invoiceDate)
	page.Click("#date1")
	for i := 0; i < 10; i++ {
		page.Keyboard().Press("Backspace")
	}
	for _, letter := range invoiceDate {
		page.Keyboard().Press(string(letter))
	}
	page.Fill("#nodate1", invoiceDate)
	page.Click("#pobierz_i_zapisz")

	eventDone := make(chan bool)
	// var wg sync.WaitGroup
	// wg.Add(1)
	page.On("download", func(download playwright.Download) {
		// download.SaveAs blocks the execution of the function - use goroutine
		go func() {
			// Save the download using download.SaveAs
			download.SaveAs(downloadFilePath)
			// fmt.Println("File downloaded: ", downloadFilePath)
			download.Delete()
			// Signal the completion of the download
			// wg.Done()
			eventDone <- true
		}()
	})

	// Wait for the 'load' event to complete
	<-eventDone
	// wg.Wait() // Wait for the event handling to complete

	pw.Stop()
}
